package service

import (
	"errors"
	"fmt"
	"sort"

	"bemunair2026/server/database/entities"
	contentService "bemunair2026/server/modules/content_submission/service"
	"bemunair2026/server/modules/letter_submission/dto"
	"bemunair2026/server/modules/letter_submission/repository"
	userRepository "bemunair2026/server/modules/user/repository"
	"bemunair2026/server/modules/wa_notification"
	"bemunair2026/server/pkg"
	"bemunair2026/server/pkg/constants"
)

type LetterSubmissionService interface {
	Create(req dto.CreateRequest, submitterID uint64, claimsMinistry *string) (*dto.LetterSubmissionResponse, []error, error)
	ListForUser(role string, userID uint64, ministry *string) ([]dto.LetterSubmissionResponse, error)
	Get(id uint64, role string, userID uint64, ministry *string) (*dto.LetterSubmissionResponse, error)
	Timeline(id uint64, role string, userID uint64, ministry *string) ([]dto.StatusHistoryResponse, error)
	SubmitRevision(id uint64, role string, userID uint64, ministry *string, note *string) (*dto.LetterSubmissionResponse, error)
	UpdateStatus(id uint64, req dto.UpdateStatusRequest, actorID uint64) (*dto.LetterSubmissionResponse, error)
	AssignPJ(id, pjID, actorID uint64) (*dto.LetterSubmissionResponse, error)
	Delete(id uint64) error
}

type letterSubmissionService struct {
	repository     repository.LetterSubmissionRepository
	userRepository userRepository.UserRepository
	wa             pkg.WASender
}

var _ LetterSubmissionService = (*letterSubmissionService)(nil)

func NewLetterSubmissionService(
	repository repository.LetterSubmissionRepository,
	userRepository userRepository.UserRepository,
	wa pkg.WASender,
) LetterSubmissionService {
	return &letterSubmissionService{repository: repository, userRepository: userRepository, wa: wa}
}

func (s *letterSubmissionService) Create(req dto.CreateRequest, submitterID uint64, claimsMinistry *string) (*dto.LetterSubmissionResponse, []error, error) {
	submitter, _ := s.userRepository.FindByID(submitterID)
	if submitter == nil {
		return nil, nil, errors.New("submitter tidak ditemukan")
	}
	if req.Ministry == "" && submitter.Ministry != nil {
		req.Ministry = *submitter.Ministry
	} else if req.Ministry == "" && claimsMinistry != nil {
		req.Ministry = *claimsMinistry
	}

	submission := &entities.LetterSubmission{
		SubmitterID:    submitterID,
		SubmitterName:  submitter.Name,
		SubmitterPhone: submitter.Phone,
		MinistryID:     submitter.MinistryID,
		Ministry:       req.Ministry,
		LetterType:     req.LetterType,
		Subject:        req.Subject,
		Body:           req.Body,
		Deadline:       req.DeadlineAt,
		Status:         constants.StatusSubmitted,
	}

	err := s.repository.Create(submission)
	if err != nil {
		return nil, nil, err
	}
	submission.Submitter = submitter
	submission.AssignedPJ = nil

	res := dto.NewLetterSubmissionResponse(submission)
	return &res, wa_notification.NotifyLetterSubmissionCreated(submission, nil, submitter, s.wa), nil
}

func (s *letterSubmissionService) ListForUser(role string, userID uint64, ministry *string) ([]dto.LetterSubmissionResponse, error) {
	rows, err := s.repository.ListForUser(role, userID, ministry)
	if err != nil {
		return nil, err
	}
	return dto.NewLetterSubmissionResponses(rows), nil
}

func (s *letterSubmissionService) Get(id uint64, role string, userID uint64, ministry *string) (*dto.LetterSubmissionResponse, error) {
	submission, err := s.repository.FindByID(id)
	if err != nil || submission == nil {
		return nil, err
	}
	if role != constants.RoleAdminMedinfo && submission.SubmitterID != userID && (ministry == nil || submission.Ministry != *ministry) {
		return nil, errors.New("forbidden")
	}
	res := dto.NewLetterSubmissionResponse(submission)
	return &res, nil
}

func (s *letterSubmissionService) Timeline(id uint64, role string, userID uint64, ministry *string) ([]dto.StatusHistoryResponse, error) {
	if _, err := s.Get(id, role, userID, ministry); err != nil {
		return nil, err
	}
	statusRows, err := s.repository.ListHistory(id)
	if err != nil {
		return nil, err
	}
	assignmentRows, err := s.repository.ListAssignmentHistory(id)
	if err != nil {
		return nil, err
	}
	events := append(dto.NewStatusHistoryResponses(statusRows), dto.NewAssignmentHistoryResponses(assignmentRows)...)
	sort.SliceStable(events, func(i, j int) bool { return events[i].CreatedAt.Before(events[j].CreatedAt) })
	return events, nil
}

func (s *letterSubmissionService) SubmitRevision(id uint64, role string, userID uint64, ministry *string, note *string) (*dto.LetterSubmissionResponse, error) {
	current, err := s.repository.FindByID(id)
	if err != nil || current == nil {
		return nil, err
	}
	if role != constants.RoleAdminMedinfo && current.SubmitterID != userID && (ministry == nil || current.Ministry != *ministry) {
		return nil, errors.New("forbidden")
	}
	if current.Status != constants.StatusRevisionRequired {
		return nil, errors.New("submission tidak sedang menunggu revisi")
	}
	updated, err := s.repository.UpdateStatus(id, constants.StatusRevisionSubmitted, note, userID)
	if err != nil {
		return nil, err
	}
	res := dto.NewLetterSubmissionResponse(updated)
	return &res, nil
}

func (s *letterSubmissionService) UpdateStatus(id uint64, req dto.UpdateStatusRequest, actorID uint64) (*dto.LetterSubmissionResponse, error) {
	current, err := s.repository.FindByID(id)
	if err != nil || current == nil {
		return nil, err
	}
	if req.Status == constants.StatusPendingReview && current.AssignedPJID == nil {
		return nil, errors.New("tetapkan PJ sebelum memulai peninjauan")
	}
	if !contentService.ValidTransition(current.Status, req.Status) && !(current.Status == constants.StatusApproved && req.Status == constants.StatusCompleted) {
		return nil, errors.New("invalid transition")
	}
	updated, err := s.repository.UpdateStatus(id, req.Status, req.Notes, actorID)
	if err != nil || updated == nil {
		return nil, err
	}
	res := dto.NewLetterSubmissionResponse(updated)
	return &res, nil
}

func (s *letterSubmissionService) AssignPJ(id, pjID, actorID uint64) (*dto.LetterSubmissionResponse, error) {
	row, err := s.repository.AssignPJ(id, pjID, actorID)
	if err != nil || row == nil {
		return nil, err
	}
	_ = wa_notification.NotifyAssignedPJ(row.AssignedPJ, fmt.Sprintf("Anda ditetapkan sebagai PJ untuk %s - %s.", pointerValue(row.RequestCode), row.Subject), s.wa)
	res := dto.NewLetterSubmissionResponse(row)
	return &res, nil
}

func pointerValue(value *string) string {
	if value == nil {
		return "pengajuan surat"
	}
	return *value
}

func (s *letterSubmissionService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

package service

import (
	"errors"

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
	Get(id uint64) (*dto.LetterSubmissionResponse, error)
	UpdateStatus(id uint64, req dto.UpdateStatusRequest) (*dto.LetterSubmissionResponse, error)
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
	if req.Ministry == "" && claimsMinistry != nil {
		req.Ministry = *claimsMinistry
	}

	submission := &entities.LetterSubmission{
		SubmitterID: submitterID,
		Ministry:    req.Ministry,
		LetterType:  req.LetterType,
		Subject:     req.Subject,
		Body:        req.Body,
		Deadline:    req.DeadlineAt,
		Status:      constants.StatusPending,
	}

	pj, err := s.repository.CreateWithAssignment(submission)
	if err != nil {
		return nil, nil, err
	}
	submission.Submitter = submitter
	submission.AssignedPJ = pj

	res := dto.NewLetterSubmissionResponse(submission)
	return &res, wa_notification.NotifyLetterSubmissionCreated(submission, pj, submitter, s.wa), nil
}

func (s *letterSubmissionService) ListForUser(role string, userID uint64, ministry *string) ([]dto.LetterSubmissionResponse, error) {
	rows, err := s.repository.ListForUser(role, userID, ministry)
	if err != nil {
		return nil, err
	}
	return dto.NewLetterSubmissionResponses(rows), nil
}

func (s *letterSubmissionService) Get(id uint64) (*dto.LetterSubmissionResponse, error) {
	submission, err := s.repository.FindByID(id)
	if err != nil || submission == nil {
		return nil, err
	}
	res := dto.NewLetterSubmissionResponse(submission)
	return &res, nil
}

func (s *letterSubmissionService) UpdateStatus(id uint64, req dto.UpdateStatusRequest) (*dto.LetterSubmissionResponse, error) {
	current, err := s.repository.FindByID(id)
	if err != nil || current == nil {
		return nil, err
	}
	if !contentService.ValidTransition(current.Status, req.Status) {
		return nil, errors.New("invalid transition")
	}
	updated, err := s.repository.UpdateStatus(id, req.Status, req.Notes)
	if err != nil || updated == nil {
		return nil, err
	}
	res := dto.NewLetterSubmissionResponse(updated)
	return &res, nil
}

func (s *letterSubmissionService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

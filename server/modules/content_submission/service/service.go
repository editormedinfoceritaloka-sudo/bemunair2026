package service

import (
	"errors"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/content_submission/dto"
	"bemunair2026/server/modules/content_submission/repository"
	userRepository "bemunair2026/server/modules/user/repository"
	"bemunair2026/server/modules/wa_notification"
	"bemunair2026/server/pkg"
	"bemunair2026/server/pkg/constants"
)

type ContentSubmissionService interface {
	Create(req dto.CreateRequest, submitterID uint64, claimsMinistry *string) (*dto.ContentSubmissionResponse, []error, error)
	ListForUser(role string, userID uint64, ministry *string) ([]dto.ContentSubmissionResponse, error)
	Get(id uint64) (*dto.ContentSubmissionResponse, error)
	UpdateStatus(id uint64, req dto.UpdateStatusRequest) (*dto.ContentSubmissionResponse, error)
	Delete(id uint64) error
}

type contentSubmissionService struct {
	repository     repository.ContentSubmissionRepository
	userRepository userRepository.UserRepository
	wa             pkg.WASender
}

var _ ContentSubmissionService = (*contentSubmissionService)(nil)

func NewContentSubmissionService(
	repository repository.ContentSubmissionRepository,
	userRepository userRepository.UserRepository,
	wa pkg.WASender,
) ContentSubmissionService {
	return &contentSubmissionService{repository: repository, userRepository: userRepository, wa: wa}
}

func (s *contentSubmissionService) Create(req dto.CreateRequest, submitterID uint64, claimsMinistry *string) (*dto.ContentSubmissionResponse, []error, error) {
	submitter, _ := s.userRepository.FindByID(submitterID)
	if req.Ministry == "" && claimsMinistry != nil {
		req.Ministry = *claimsMinistry
	}

	submission := &entities.ContentSubmission{
		SubmitterID:    submitterID,
		Ministry:       req.Ministry,
		Platform:       req.Platform,
		SubmissionType: req.SubmissionType,
		Caption:        req.Caption,
		Deadline:       req.Deadline,
		BriefFile:      req.BriefFile,
		PosterFile:     req.PosterFile,
		Status:         constants.StatusPending,
	}

	pj, err := s.repository.CreateWithAssignment(submission)
	if err != nil {
		return nil, nil, err
	}
	submission.Submitter = submitter
	submission.AssignedPJ = pj

	res := dto.NewContentSubmissionResponse(submission)
	return &res, wa_notification.NotifyContentSubmissionCreated(submission, pj, submitter, s.wa), nil
}

func (s *contentSubmissionService) ListForUser(role string, userID uint64, ministry *string) ([]dto.ContentSubmissionResponse, error) {
	rows, err := s.repository.ListForUser(role, userID, ministry)
	if err != nil {
		return nil, err
	}
	return dto.NewContentSubmissionResponses(rows), nil
}

func (s *contentSubmissionService) Get(id uint64) (*dto.ContentSubmissionResponse, error) {
	submission, err := s.repository.FindByID(id)
	if err != nil || submission == nil {
		return nil, err
	}
	res := dto.NewContentSubmissionResponse(submission)
	return &res, nil
}

func (s *contentSubmissionService) UpdateStatus(id uint64, req dto.UpdateStatusRequest) (*dto.ContentSubmissionResponse, error) {
	current, err := s.repository.FindByID(id)
	if err != nil || current == nil {
		return nil, err
	}
	if !ValidTransition(current.Status, req.Status) {
		return nil, errors.New("invalid transition")
	}
	updated, err := s.repository.UpdateStatus(id, req.Status, req.Notes)
	if err != nil || updated == nil {
		return nil, err
	}
	res := dto.NewContentSubmissionResponse(updated)
	return &res, nil
}

func (s *contentSubmissionService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func ValidTransition(from, to string) bool {
	if from == constants.StatusPending && to == constants.StatusInReview {
		return true
	}
	if from == constants.StatusInReview && (to == constants.StatusApproved || to == constants.StatusRejected) {
		return true
	}
	return from == to
}

package service

import (
	"errors"
	"time"

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
	Get(id uint64, role string, userID uint64, ministry *string) (*dto.ContentSubmissionResponse, error)
	Timeline(id uint64, role string, userID uint64, ministry *string) ([]dto.StatusHistoryResponse, error)
	UpdateStatus(id uint64, req dto.UpdateStatusRequest, actorID uint64) (*dto.ContentSubmissionResponse, error)
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
	if submitter == nil {
		return nil, nil, errors.New("submitter tidak ditemukan")
	}
	if req.Ministry == "" && submitter.Ministry != nil {
		req.Ministry = *submitter.Ministry
	} else if req.Ministry == "" && claimsMinistry != nil {
		req.Ministry = *claimsMinistry
	}
	if req.MinistryID == nil {
		req.MinistryID = submitter.MinistryID
	}
	if req.SubmitterPhone == nil {
		req.SubmitterPhone = submitter.Phone
	}
	if req.ServiceType == "" {
		if req.SubmissionType == constants.ContentTypeArtikel {
			req.ServiceType = constants.ServiceTypeArticle
		} else {
			req.ServiceType = constants.ServiceTypeContent
		}
	}

	deadline := deriveDeadline(req)

	submission := &entities.ContentSubmission{
		SubmitterID:            submitterID,
		SubmitterName:          submitter.Name,
		SubmitterPhone:         req.SubmitterPhone,
		MinistryID:             req.MinistryID,
		ServiceType:            req.ServiceType,
		ContentFormat:          req.ContentFormat,
		Ministry:               req.Ministry,
		SubmissionType:         req.SubmissionType,
		Title:                  req.Title,
		AddSong:                req.AddSong,
		SongTitle:              req.SongTitle,
		SongArtist:             req.SongArtist,
		SongStartSeconds:       req.SongStartSeconds,
		SongEndSeconds:         req.SongEndSeconds,
		Caption:                req.Caption,
		AdditionalNotes:        req.AdditionalNotes,
		PublishDate:            req.PublishDate,
		PublishTime:            req.PublishTime,
		DesignDriveLink:        req.DesignDriveLink,
		CanvaLink:              req.CanvaLink,
		ArticleDriveLink:       req.ArticleDriveLink,
		DocumentationDriveLink: req.DocumentationDriveLink,
		RequiredInformation:    req.RequiredInformation,
		Deadline:               deadline,
		BriefLink:              req.BriefLink,
		Status:                 constants.StatusSubmitted,
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

func (s *contentSubmissionService) Get(id uint64, role string, userID uint64, ministry *string) (*dto.ContentSubmissionResponse, error) {
	submission, err := s.repository.FindByID(id)
	if err != nil || submission == nil {
		return nil, err
	}
	if role != constants.RoleAdminMedinfo && submission.SubmitterID != userID && (ministry == nil || submission.Ministry != *ministry) {
		return nil, errors.New("forbidden")
	}
	res := dto.NewContentSubmissionResponse(submission)
	return &res, nil
}

func (s *contentSubmissionService) Timeline(id uint64, role string, userID uint64, ministry *string) ([]dto.StatusHistoryResponse, error) {
	if _, err := s.Get(id, role, userID, ministry); err != nil {
		return nil, err
	}
	rows, err := s.repository.ListHistory(id)
	if err != nil {
		return nil, err
	}
	return dto.NewStatusHistoryResponses(rows), nil
}

func (s *contentSubmissionService) UpdateStatus(id uint64, req dto.UpdateStatusRequest, actorID uint64) (*dto.ContentSubmissionResponse, error) {
	current, err := s.repository.FindByID(id)
	if err != nil || current == nil {
		return nil, err
	}
	if !ValidTransition(current.Status, req.Status) {
		return nil, errors.New("invalid transition")
	}
	updated, err := s.repository.UpdateStatus(id, req.Status, req.Notes, actorID)
	if err != nil || updated == nil {
		return nil, err
	}
	res := dto.NewContentSubmissionResponse(updated)
	return &res, nil
}

func (s *contentSubmissionService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func deriveDeadline(req dto.CreateRequest) *time.Time {
	if req.PublishDate == nil {
		return nil
	}
	date := *req.PublishDate
	hour, minute := 0, 0
	if req.PublishTime != nil {
		if t, err := time.Parse("15:04", *req.PublishTime); err == nil {
			hour, minute = t.Hour(), t.Minute()
		}
	}
	deadline := time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, time.Local)
	return &deadline
}

func ValidTransition(from, to string) bool {
	allowed := map[string]map[string]bool{
		constants.StatusPending:           {constants.StatusInReview: true},
		constants.StatusInReview:          {constants.StatusApproved: true, constants.StatusRejected: true},
		constants.StatusSubmitted:         {constants.StatusPendingReview: true, constants.StatusRejected: true},
		constants.StatusPendingReview:     {constants.StatusRevisionRequired: true, constants.StatusApproved: true, constants.StatusRejected: true},
		constants.StatusRevisionSubmitted: {constants.StatusPendingReview: true, constants.StatusRevisionRequired: true, constants.StatusApproved: true, constants.StatusRejected: true},
		constants.StatusApproved:          {constants.StatusScheduled: true, constants.StatusRejected: true},
		constants.StatusScheduled:         {constants.StatusPublished: true},
	}
	return from == to || allowed[from][to]
}

package content_submission

import (
	"errors"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/wa_notification"
	"bemunair2026/server/pkg"
	"bemunair2026/server/pkg/constants"
)

type Service struct {
	repo *Repository
	wa   pkg.WASender
}

func NewService(repo *Repository, wa pkg.WASender) *Service { return &Service{repo: repo, wa: wa} }

func (s *Service) Create(sub *entities.ContentSubmission, submitter *entities.User) (*entities.ContentSubmission, []error, error) {
	pj, err := s.repo.CreateWithAssignment(sub)
	if err != nil {
		return nil, nil, err
	}
	sub.Submitter = submitter
	sub.AssignedPJ = pj
	return sub, wa_notification.NotifyContentSubmissionCreated(sub, pj, submitter, s.wa), nil
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

func (s *Service) UpdateStatus(id uint64, status string, notes *string) (*entities.ContentSubmission, error) {
	current, err := s.repo.FindByID(id)
	if err != nil || current == nil {
		return current, err
	}
	if !ValidTransition(current.Status, status) {
		return nil, errors.New("invalid transition")
	}
	return s.repo.UpdateStatus(id, status, notes)
}

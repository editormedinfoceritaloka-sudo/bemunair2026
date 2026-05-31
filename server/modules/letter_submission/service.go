package letter_submission

import (
	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/content_submission"
	"bemunair2026/server/modules/wa_notification"
	"bemunair2026/server/pkg"
	"errors"
)

type Service struct {
	repo *Repository
	wa   pkg.WASender
}

func NewService(repo *Repository, wa pkg.WASender) *Service { return &Service{repo: repo, wa: wa} }

func (s *Service) Create(sub *entities.LetterSubmission, submitter *entities.User) (*entities.LetterSubmission, []error, error) {
	pj, err := s.repo.CreateWithAssignment(sub)
	if err != nil {
		return nil, nil, err
	}
	sub.Submitter = submitter
	sub.AssignedPJ = pj
	return sub, wa_notification.NotifyLetterSubmissionCreated(sub, pj, submitter, s.wa), nil
}

func (s *Service) UpdateStatus(id uint64, status string, notes *string) (*entities.LetterSubmission, error) {
	current, err := s.repo.FindByID(id)
	if err != nil || current == nil {
		return current, err
	}
	if !content_submission.ValidTransition(current.Status, status) {
		return nil, errors.New("invalid transition")
	}
	return s.repo.UpdateStatus(id, status, notes)
}

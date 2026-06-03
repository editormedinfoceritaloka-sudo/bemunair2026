package repository

import (
	"errors"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/medinfo_pj"
	"bemunair2026/server/pkg/constants"
	"gorm.io/gorm"
)

type LetterSubmissionRepository interface {
	CreateWithAssignment(submission *entities.LetterSubmission) (*entities.User, error)
	FindByID(id uint64) (*entities.LetterSubmission, error)
	ListForUser(role string, userID uint64, ministry *string) ([]entities.LetterSubmission, error)
	UpdateStatus(id uint64, status string, notes *string) (*entities.LetterSubmission, error)
	Delete(id uint64) error
	ListPendingOlderThan(age time.Duration) ([]entities.LetterSubmission, error)
}

type letterSubmissionRepository struct {
	db *gorm.DB
}

var _ LetterSubmissionRepository = (*letterSubmissionRepository)(nil)

func NewLetterSubmissionRepository(db *gorm.DB) LetterSubmissionRepository {
	return &letterSubmissionRepository{db: db}
}

func (r *letterSubmissionRepository) CreateWithAssignment(submission *entities.LetterSubmission) (*entities.User, error) {
	var pj *entities.User
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if submission.Ministry == "MEDINFO" {
			assigned, err := medinfo_pj.AssignNext(tx)
			if err != nil {
				return err
			}
			pj = assigned
			if pj != nil {
				submission.AssignedPJID = &pj.ID
			}
		}
		return tx.Create(submission).Error
	})
	return pj, err
}

func (r *letterSubmissionRepository) FindByID(id uint64) (*entities.LetterSubmission, error) {
	var submission entities.LetterSubmission
	err := r.db.Preload("Submitter").Preload("AssignedPJ").First(&submission, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &submission, err
}

func (r *letterSubmissionRepository) ListForUser(role string, userID uint64, ministry *string) ([]entities.LetterSubmission, error) {
	var rows []entities.LetterSubmission
	query := r.db.Preload("Submitter").Preload("AssignedPJ").Order("deadline ASC")
	if role == constants.RoleMentri {
		query = query.Where("submitter_id = ? OR ministry = ?", userID, value(ministry))
	}
	return rows, query.Find(&rows).Error
}

func (r *letterSubmissionRepository) UpdateStatus(id uint64, status string, notes *string) (*entities.LetterSubmission, error) {
	submission, err := r.FindByID(id)
	if err != nil || submission == nil {
		return submission, err
	}
	submission.Status = status
	submission.Notes = notes
	return submission, r.db.Save(submission).Error
}

func (r *letterSubmissionRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.LetterSubmission{}, id).Error
}

func (r *letterSubmissionRepository) ListPendingOlderThan(age time.Duration) ([]entities.LetterSubmission, error) {
	var rows []entities.LetterSubmission
	return rows, r.db.Preload("Submitter").Preload("AssignedPJ").
		Where("status IN ? AND created_at <= ?", []string{constants.StatusPending, constants.StatusInReview}, time.Now().Add(-age)).
		Order("deadline ASC").
		Find(&rows).Error
}

func value(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

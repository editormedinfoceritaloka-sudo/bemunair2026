package repository

import (
	"errors"
	"time"

	"bemunair2026/server/database/entities"
	medinfoRepository "bemunair2026/server/modules/medinfo_pj/repository"
	"bemunair2026/server/pkg/constants"
	"gorm.io/gorm"
)

type ContentSubmissionRepository interface {
	CreateWithAssignment(submission *entities.ContentSubmission) (*entities.User, error)
	FindByID(id uint64) (*entities.ContentSubmission, error)
	ListForUser(role string, userID uint64, ministry *string) ([]entities.ContentSubmission, error)
	UpdateStatus(id uint64, status string, notes *string) (*entities.ContentSubmission, error)
	Delete(id uint64) error
	ListPendingOlderThan(age time.Duration) ([]entities.ContentSubmission, error)
}

type contentSubmissionRepository struct {
	db *gorm.DB
}

var _ ContentSubmissionRepository = (*contentSubmissionRepository)(nil)

func NewContentSubmissionRepository(db *gorm.DB) ContentSubmissionRepository {
	return &contentSubmissionRepository{db: db}
}

func (r *contentSubmissionRepository) CreateWithAssignment(submission *entities.ContentSubmission) (*entities.User, error) {
	var pj *entities.User
	err := r.db.Transaction(func(tx *gorm.DB) error {
		assigned, err := medinfoRepository.AssignNext(tx)
		if err != nil {
			return err
		}
		pj = assigned
		if pj != nil {
			submission.AssignedPJID = &pj.ID
		}
		return tx.Create(submission).Error
	})
	return pj, err
}

func (r *contentSubmissionRepository) FindByID(id uint64) (*entities.ContentSubmission, error) {
	var submission entities.ContentSubmission
	err := r.db.Preload("Submitter").Preload("AssignedPJ").First(&submission, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &submission, err
}

func (r *contentSubmissionRepository) ListForUser(role string, userID uint64, ministry *string) ([]entities.ContentSubmission, error) {
	var rows []entities.ContentSubmission
	query := r.db.Preload("Submitter").Preload("AssignedPJ").Order("deadline IS NULL, deadline ASC")
	if role == constants.RoleMentri {
		query = query.Where("submitter_id = ? OR ministry = ?", userID, value(ministry))
	}
	return rows, query.Find(&rows).Error
}

func (r *contentSubmissionRepository) UpdateStatus(id uint64, status string, notes *string) (*entities.ContentSubmission, error) {
	submission, err := r.FindByID(id)
	if err != nil || submission == nil {
		return submission, err
	}
	submission.Status = status
	submission.Notes = notes
	return submission, r.db.Save(submission).Error
}

func (r *contentSubmissionRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.ContentSubmission{}, id).Error
}

func (r *contentSubmissionRepository) ListPendingOlderThan(age time.Duration) ([]entities.ContentSubmission, error) {
	var rows []entities.ContentSubmission
	cutoff := time.Now().Add(-age)
	return rows, r.db.Preload("Submitter").Preload("AssignedPJ").
		Where("status IN ? AND created_at <= ? AND deadline IS NOT NULL", []string{constants.StatusPending, constants.StatusInReview}, cutoff).
		Order("deadline ASC").
		Find(&rows).Error
}

func value(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

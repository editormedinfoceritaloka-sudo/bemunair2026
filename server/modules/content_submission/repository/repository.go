package repository

import (
	"errors"
	"fmt"
	"time"

	"bemunair2026/server/database/entities"
	medinfoRepository "bemunair2026/server/modules/medinfo_pj/repository"
	"bemunair2026/server/pkg/constants"
	"gorm.io/gorm"
)

type ContentSubmissionRepository interface {
	CreateWithAssignment(submission *entities.ContentSubmission) (*entities.User, error)
	FindByID(id uint64) (*entities.ContentSubmission, error)
	ListHistory(id uint64) ([]entities.ContentSubmissionStatusHistory, error)
	ListForUser(role string, userID uint64, ministry *string) ([]entities.ContentSubmission, error)
	UpdateStatus(id uint64, status string, notes *string, actorID uint64) (*entities.ContentSubmission, error)
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
		if err := tx.Create(submission).Error; err != nil {
			return err
		}
		code := fmt.Sprintf("MED-%d-%06d", submission.CreatedAt.Year(), submission.ID)
		submission.RequestCode = &code
		if err := tx.Model(submission).Update("request_code", code).Error; err != nil {
			return err
		}
		return tx.Create(&entities.ContentSubmissionStatusHistory{SubmissionID: submission.ID, ActorID: &submission.SubmitterID, ToStatus: submission.Status, Note: stringPointer("Pengajuan dikirim")}).Error
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

func (r *contentSubmissionRepository) ListHistory(id uint64) ([]entities.ContentSubmissionStatusHistory, error) {
	var rows []entities.ContentSubmissionStatusHistory
	return rows, r.db.Preload("Actor").Where("submission_id = ?", id).Order("created_at ASC").Find(&rows).Error
}

func (r *contentSubmissionRepository) ListForUser(role string, userID uint64, ministry *string) ([]entities.ContentSubmission, error) {
	var rows []entities.ContentSubmission
	query := r.db.Preload("Submitter").Preload("AssignedPJ").Order("deadline IS NULL, deadline ASC")
	if role != constants.RoleAdminMedinfo {
		query = query.Where("submitter_id = ? OR ministry = ?", userID, value(ministry))
	}
	return rows, query.Find(&rows).Error
}

func (r *contentSubmissionRepository) UpdateStatus(id uint64, status string, notes *string, actorID uint64) (*entities.ContentSubmission, error) {
	submission, err := r.FindByID(id)
	if err != nil || submission == nil {
		return submission, err
	}
	from := submission.Status
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(submission).Updates(map[string]any{"status": status, "notes": notes}).Error; err != nil {
			return err
		}
		return tx.Create(&entities.ContentSubmissionStatusHistory{SubmissionID: id, ActorID: &actorID, FromStatus: &from, ToStatus: status, Note: notes}).Error
	})
	if err == nil {
		submission.Status = status
		submission.Notes = notes
	}
	return submission, err
}

func (r *contentSubmissionRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.ContentSubmission{}, id).Error
}

func (r *contentSubmissionRepository) ListPendingOlderThan(age time.Duration) ([]entities.ContentSubmission, error) {
	var rows []entities.ContentSubmission
	cutoff := time.Now().Add(-age)
	return rows, r.db.Preload("Submitter").Preload("AssignedPJ").
		Where("status IN ? AND created_at <= ? AND deadline IS NOT NULL", []string{constants.StatusSubmitted, constants.StatusPendingReview, constants.StatusRevisionSubmitted}, cutoff).
		Order("deadline ASC").
		Find(&rows).Error
}

func stringPointer(value string) *string { return &value }

func value(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

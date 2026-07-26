package repository

import (
	"errors"
	"fmt"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/pkg/constants"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContentSubmissionRepository interface {
	Create(submission *entities.ContentSubmission) error
	FindByID(id uint64) (*entities.ContentSubmission, error)
	ListHistory(id uint64) ([]entities.ContentSubmissionStatusHistory, error)
	ListAssignmentHistory(id uint64) ([]entities.ContentSubmissionAssignmentHistory, error)
	ListForUser(role string, userID uint64, ministry *string) ([]entities.ContentSubmission, error)
	UpdateStatus(id uint64, status string, notes *string, actorID uint64) (*entities.ContentSubmission, error)
	AssignPJ(id, pjID, actorID uint64) (*entities.ContentSubmission, error)
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

func (r *contentSubmissionRepository) Create(submission *entities.ContentSubmission) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
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
}

func (r *contentSubmissionRepository) FindByID(id uint64) (*entities.ContentSubmission, error) {
	var submission entities.ContentSubmission
	err := r.db.Preload("Submitter").Preload("AssignedPJ").Preload("Attachments").First(&submission, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &submission, err
}

func (r *contentSubmissionRepository) ListHistory(id uint64) ([]entities.ContentSubmissionStatusHistory, error) {
	var rows []entities.ContentSubmissionStatusHistory
	return rows, r.db.Preload("Actor").Where("submission_id = ?", id).Order("created_at ASC").Find(&rows).Error
}

func (r *contentSubmissionRepository) ListAssignmentHistory(id uint64) ([]entities.ContentSubmissionAssignmentHistory, error) {
	var rows []entities.ContentSubmissionAssignmentHistory
	return rows, r.db.Preload("Actor").Preload("FromPJ").Preload("ToPJ").Where("submission_id = ?", id).Order("created_at ASC").Find(&rows).Error
}

func (r *contentSubmissionRepository) ListForUser(role string, userID uint64, ministry *string) ([]entities.ContentSubmission, error) {
	var rows []entities.ContentSubmission
	query := r.db.Preload("Submitter").Preload("AssignedPJ").Preload("Attachments").Order("deadline IS NULL, deadline ASC")
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

func (r *contentSubmissionRepository) AssignPJ(id, pjID, actorID uint64) (*entities.ContentSubmission, error) {
	var submission entities.ContentSubmission
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var queue entities.MedinfoPJQueue
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ?", pjID).First(&queue).Error; err != nil {
			return errors.New("PJ tidak terdaftar pada roster Medinfo")
		}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&submission, id).Error; err != nil {
			return err
		}
		activeContent := []string{constants.StatusSubmitted, constants.StatusPendingReview, constants.StatusRevisionRequired, constants.StatusRevisionSubmitted, constants.StatusApproved, constants.StatusScheduled}
		activeLetter := []string{constants.StatusSubmitted, constants.StatusPendingReview, constants.StatusRevisionRequired, constants.StatusRevisionSubmitted, constants.StatusApproved}
		assignable := false
		for _, status := range activeContent {
			if submission.Status == status {
				assignable = true
				break
			}
		}
		if !assignable {
			return errors.New("task yang sudah selesai tidak dapat di-assign")
		}
		var busy int64
		if err := tx.Model(&entities.ContentSubmission{}).Where("assigned_pj_id = ? AND id <> ? AND status IN ?", pjID, id, activeContent).Count(&busy).Error; err != nil {
			return err
		}
		if busy == 0 {
			if err := tx.Model(&entities.LetterSubmission{}).Where("assigned_pj_id = ? AND status IN ?", pjID, activeLetter).Count(&busy).Error; err != nil {
				return err
			}
		}
		if busy > 0 {
			return errors.New("PJ sedang menangani task aktif")
		}
		from := submission.AssignedPJID
		if from != nil && *from == pjID {
			return nil
		}
		if err := tx.Model(&submission).Update("assigned_pj_id", pjID).Error; err != nil {
			return err
		}
		note := "PJ Medinfo ditetapkan"
		return tx.Create(&entities.ContentSubmissionAssignmentHistory{SubmissionID: id, ActorID: actorID, FromPJID: from, ToPJID: pjID, Note: &note}).Error
	})
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *contentSubmissionRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.ContentSubmission{}, id).Error
}

func (r *contentSubmissionRepository) ListPendingOlderThan(age time.Duration) ([]entities.ContentSubmission, error) {
	var rows []entities.ContentSubmission
	cutoff := time.Now().Add(-age)
	return rows, r.db.Preload("Submitter").Preload("AssignedPJ").Preload("Attachments").
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

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

type LetterSubmissionRepository interface {
	Create(submission *entities.LetterSubmission) error
	FindByID(id uint64) (*entities.LetterSubmission, error)
	ListHistory(id uint64) ([]entities.LetterSubmissionStatusHistory, error)
	ListAssignmentHistory(id uint64) ([]entities.LetterSubmissionAssignmentHistory, error)
	ListForUser(role string, userID uint64, ministry *string) ([]entities.LetterSubmission, error)
	UpdateStatus(id uint64, status string, notes *string, actorID uint64) (*entities.LetterSubmission, error)
	AssignPJ(id, pjID, actorID uint64) (*entities.LetterSubmission, error)
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

func (r *letterSubmissionRepository) Create(submission *entities.LetterSubmission) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(submission).Error; err != nil {
			return err
		}
		code := fmt.Sprintf("SUR-%d-%06d", submission.CreatedAt.Year(), submission.ID)
		submission.RequestCode = &code
		if err := tx.Model(submission).Update("request_code", code).Error; err != nil {
			return err
		}
		note := "Pengajuan surat dikirim"
		return tx.Create(&entities.LetterSubmissionStatusHistory{SubmissionID: submission.ID, ActorID: &submission.SubmitterID, ToStatus: submission.Status, Note: &note}).Error
	})
}

func (r *letterSubmissionRepository) FindByID(id uint64) (*entities.LetterSubmission, error) {
	var submission entities.LetterSubmission
	err := r.db.Preload("Submitter").Preload("AssignedPJ").First(&submission, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &submission, err
}

func (r *letterSubmissionRepository) ListHistory(id uint64) ([]entities.LetterSubmissionStatusHistory, error) {
	var rows []entities.LetterSubmissionStatusHistory
	return rows, r.db.Preload("Actor").Where("submission_id = ?", id).Order("created_at ASC").Find(&rows).Error
}

func (r *letterSubmissionRepository) ListAssignmentHistory(id uint64) ([]entities.LetterSubmissionAssignmentHistory, error) {
	var rows []entities.LetterSubmissionAssignmentHistory
	return rows, r.db.Preload("Actor").Preload("FromPJ").Preload("ToPJ").Where("submission_id = ?", id).Order("created_at ASC").Find(&rows).Error
}

func (r *letterSubmissionRepository) ListForUser(role string, userID uint64, ministry *string) ([]entities.LetterSubmission, error) {
	var rows []entities.LetterSubmission
	query := r.db.Preload("Submitter").Preload("AssignedPJ").Order("deadline ASC")
	if role != constants.RoleAdminMedinfo {
		query = query.Where("submitter_id = ? OR ministry = ?", userID, value(ministry))
	}
	return rows, query.Find(&rows).Error
}

func (r *letterSubmissionRepository) UpdateStatus(id uint64, status string, notes *string, actorID uint64) (*entities.LetterSubmission, error) {
	submission, err := r.FindByID(id)
	if err != nil || submission == nil {
		return submission, err
	}
	from := submission.Status
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(submission).Updates(map[string]any{"status": status, "notes": notes}).Error; err != nil {
			return err
		}
		return tx.Create(&entities.LetterSubmissionStatusHistory{SubmissionID: id, ActorID: &actorID, FromStatus: &from, ToStatus: status, Note: notes}).Error
	})
	if err == nil {
		submission.Status = status
		submission.Notes = notes
	}
	return submission, err
}

func (r *letterSubmissionRepository) AssignPJ(id, pjID, actorID uint64) (*entities.LetterSubmission, error) {
	var submission entities.LetterSubmission
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
		for _, status := range activeLetter {
			if submission.Status == status {
				assignable = true
				break
			}
		}
		if !assignable {
			return errors.New("task yang sudah selesai tidak dapat di-assign")
		}
		var busy int64
		if err := tx.Model(&entities.ContentSubmission{}).Where("assigned_pj_id = ? AND status IN ?", pjID, activeContent).Count(&busy).Error; err != nil {
			return err
		}
		if busy == 0 {
			if err := tx.Model(&entities.LetterSubmission{}).Where("assigned_pj_id = ? AND id <> ? AND status IN ?", pjID, id, activeLetter).Count(&busy).Error; err != nil {
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
		return tx.Create(&entities.LetterSubmissionAssignmentHistory{SubmissionID: id, ActorID: actorID, FromPJID: from, ToPJID: pjID, Note: &note}).Error
	})
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *letterSubmissionRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.LetterSubmission{}, id).Error
}

func (r *letterSubmissionRepository) ListPendingOlderThan(age time.Duration) ([]entities.LetterSubmission, error) {
	var rows []entities.LetterSubmission
	return rows, r.db.Preload("Submitter").Preload("AssignedPJ").
		Where("status IN ? AND created_at <= ?", []string{constants.StatusSubmitted, constants.StatusPendingReview, constants.StatusRevisionSubmitted}, time.Now().Add(-age)).
		Order("deadline ASC").
		Find(&rows).Error
}

func value(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

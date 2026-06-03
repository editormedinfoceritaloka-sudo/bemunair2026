package letter_submission

import (
	"errors"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/medinfo_pj"
	"bemunair2026/server/pkg/constants"
	"gorm.io/gorm"
)

type Repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) CreateWithAssignment(s *entities.LetterSubmission) (*entities.User, error) {
	var pj *entities.User
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if s.Ministry == "MEDINFO" {
			assigned, err := medinfo_pj.AssignNext(tx)
			if err != nil {
				return err
			}
			pj = assigned
			if pj != nil {
				s.AssignedPJID = &pj.ID
			}
		}
		return tx.Create(s).Error
	})
	return pj, err
}

func (r *Repository) FindByID(id uint64) (*entities.LetterSubmission, error) {
	var s entities.LetterSubmission
	err := r.db.Preload("Submitter").Preload("AssignedPJ").First(&s, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &s, err
}

func (r *Repository) ListForUser(role string, userID uint64, ministry *string) ([]entities.LetterSubmission, error) {
	var rows []entities.LetterSubmission
	q := r.db.Preload("Submitter").Preload("AssignedPJ").Order("deadline ASC")
	if role == constants.RoleMentri {
		q = q.Where("submitter_id = ? OR ministry = ?", userID, value(ministry))
	}
	return rows, q.Find(&rows).Error
}

func (r *Repository) UpdateStatus(id uint64, status string, notes *string) (*entities.LetterSubmission, error) {
	s, err := r.FindByID(id)
	if err != nil || s == nil {
		return s, err
	}
	s.Status, s.Notes = status, notes
	return s, r.db.Save(s).Error
}

func (r *Repository) Delete(id uint64) error {
	return r.db.Delete(&entities.LetterSubmission{}, id).Error
}

func (r *Repository) ListPendingOlderThan(age time.Duration) ([]entities.LetterSubmission, error) {
	var rows []entities.LetterSubmission
	return rows, r.db.Preload("Submitter").Preload("AssignedPJ").Where("status IN ? AND created_at <= ?", []string{constants.StatusPending, constants.StatusInReview}, time.Now().Add(-age)).Order("deadline ASC").Find(&rows).Error
}

func value(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

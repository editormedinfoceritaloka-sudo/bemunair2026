package repository

import (
	"bemunair2026/server/database/entities"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QueueAvailability struct {
	Queue             entities.MedinfoPJQueue
	IsBusy            bool
	ActiveTaskType    string
	ActiveTaskID      uint64
	ActiveRequestCode string
	ActiveTaskTitle   string
}

type MedinfoPJRepository interface {
	List() ([]entities.MedinfoPJQueue, error)
	ListAvailability() ([]QueueAvailability, error)
	Create(row *entities.MedinfoPJQueue) error
	Delete(id uint64) error
	Reorder(ids []uint64) error
}

type medinfoPJRepository struct {
	db *gorm.DB
}

var _ MedinfoPJRepository = (*medinfoPJRepository)(nil)

func NewMedinfoPJRepository(db *gorm.DB) MedinfoPJRepository {
	return &medinfoPJRepository{db: db}
}

func (r *medinfoPJRepository) List() ([]entities.MedinfoPJQueue, error) {
	var rows []entities.MedinfoPJQueue
	return rows, r.db.Preload("User").Order("position ASC").Find(&rows).Error
}

func (r *medinfoPJRepository) ListAvailability() ([]QueueAvailability, error) {
	rows, err := r.List()
	if err != nil {
		return nil, err
	}
	result := make([]QueueAvailability, 0, len(rows))
	activeContent := []string{"SUBMITTED", "PENDING_REVIEW", "REVISION_REQUIRED", "REVISION_SUBMITTED", "APPROVED", "SCHEDULED"}
	activeLetter := []string{"SUBMITTED", "PENDING_REVIEW", "REVISION_REQUIRED", "REVISION_SUBMITTED", "APPROVED"}
	for i := range rows {
		item := QueueAvailability{Queue: rows[i]}
		var content entities.ContentSubmission
		err := r.db.Where("assigned_pj_id = ? AND status IN ?", rows[i].UserID, activeContent).Order("created_at ASC").First(&content).Error
		if err == nil {
			item.IsBusy = true
			item.ActiveTaskType = "CONTENT"
			item.ActiveTaskID = content.ID
			item.ActiveTaskTitle = content.Title
			if content.RequestCode != nil {
				item.ActiveRequestCode = *content.RequestCode
			}
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if !item.IsBusy {
			var letter entities.LetterSubmission
			err = r.db.Where("assigned_pj_id = ? AND status IN ?", rows[i].UserID, activeLetter).Order("created_at ASC").First(&letter).Error
			if err == nil {
				item.IsBusy = true
				item.ActiveTaskType = "LETTER"
				item.ActiveTaskID = letter.ID
				item.ActiveTaskTitle = letter.Subject
				if letter.RequestCode != nil {
					item.ActiveRequestCode = *letter.RequestCode
				}
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}
		result = append(result, item)
	}
	return result, nil
}

func (r *medinfoPJRepository) Create(row *entities.MedinfoPJQueue) error {
	var user entities.User
	if err := r.db.Preload("MinistryRef").First(&user, row.UserID).Error; err != nil {
		return err
	}
	if user.Role != "ADMIN_MEDINFO" || user.MinistryRef == nil || user.MinistryRef.Code != "MEDINFO" {
		return errors.New("PJ harus ADMIN_MEDINFO dari kementerian MEDINFO")
	}
	return r.db.Create(row).Error
}

func (r *medinfoPJRepository) Delete(id uint64) error {
	rows, err := r.ListAvailability()
	if err != nil {
		return err
	}
	for i := range rows {
		if rows[i].Queue.ID == id && rows[i].IsBusy {
			return errors.New("PJ yang sedang menangani task tidak dapat dihapus")
		}
	}
	return r.db.Delete(&entities.MedinfoPJQueue{}, id).Error
}

func (r *medinfoPJRepository) Reorder(ids []uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for i, id := range ids {
			if err := tx.Model(&entities.MedinfoPJQueue{}).Where("id = ?", id).Updates(map[string]any{"position": i + 1, "is_current": i == 0}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func AssignNext(tx *gorm.DB) (*entities.User, error) {
	var rows []entities.MedinfoPJQueue
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Preload("User").Order("position ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	idx := 0
	for i, row := range rows {
		if row.IsCurrent {
			idx = i
			break
		}
	}
	selected := rows[idx]
	next := rows[(idx+1)%len(rows)]
	if err := tx.Model(&entities.MedinfoPJQueue{}).Where("id IN ?", queueIDs(rows)).Update("is_current", false).Error; err != nil {
		return nil, err
	}
	if err := tx.Model(&entities.MedinfoPJQueue{}).Where("id = ?", next.ID).Update("is_current", true).Error; err != nil {
		return nil, err
	}
	return selected.User, nil
}

func queueIDs(rows []entities.MedinfoPJQueue) []uint64 {
	ids := make([]uint64, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.ID)
	}
	return ids
}

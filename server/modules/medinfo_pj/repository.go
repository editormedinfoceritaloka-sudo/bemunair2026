package medinfo_pj

import (
	"bemunair2026/server/database/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) List() ([]entities.MedinfoPJQueue, error) {
	var rows []entities.MedinfoPJQueue
	return rows, r.db.Preload("User").Order("position ASC").Find(&rows).Error
}

func (r *Repository) Create(row *entities.MedinfoPJQueue) error { return r.db.Create(row).Error }
func (r *Repository) Delete(id uint64) error {
	return r.db.Delete(&entities.MedinfoPJQueue{}, id).Error
}

func (r *Repository) Reorder(ids []uint64) error {
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

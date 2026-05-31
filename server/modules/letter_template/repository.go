package letter_template

import (
	"bemunair2026/server/database/entities"
	"errors"
	"gorm.io/gorm"
)

type Repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) *Repository                   { return &Repository{db: db} }
func (r *Repository) Create(t *entities.LetterTemplate) error { return r.db.Create(t).Error }
func (r *Repository) List() ([]entities.LetterTemplate, error) {
	var rows []entities.LetterTemplate
	return rows, r.db.Order("id ASC").Find(&rows).Error
}
func (r *Repository) FindByID(id uint64) (*entities.LetterTemplate, error) {
	var t entities.LetterTemplate
	err := r.db.First(&t, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &t, err
}
func (r *Repository) Update(t *entities.LetterTemplate) error { return r.db.Save(t).Error }
func (r *Repository) Delete(id uint64) error {
	return r.db.Delete(&entities.LetterTemplate{}, id).Error
}

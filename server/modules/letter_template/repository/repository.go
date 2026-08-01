package repository

import (
	"errors"

	"bemunair2026/server/database/entities"
	"gorm.io/gorm"
)

type LetterTemplateRepository interface {
	Create(template *entities.LetterTemplate) error
	List() ([]entities.LetterTemplate, error)
	FindByID(id uint64) (*entities.LetterTemplate, error)
	Update(template *entities.LetterTemplate) error
	Delete(id uint64) error
	FindMediaAsset(id uint64) (*entities.MediaAsset, error)
}

type letterTemplateRepository struct {
	db *gorm.DB
}

var _ LetterTemplateRepository = (*letterTemplateRepository)(nil)

func NewLetterTemplateRepository(db *gorm.DB) LetterTemplateRepository {
	return &letterTemplateRepository{db: db}
}

func (r *letterTemplateRepository) Create(template *entities.LetterTemplate) error {
	return r.db.Create(template).Error
}

func (r *letterTemplateRepository) List() ([]entities.LetterTemplate, error) {
	var rows []entities.LetterTemplate
	return rows, r.db.Preload("MediaAsset").Order("display_order ASC, id ASC").Find(&rows).Error
}

func (r *letterTemplateRepository) FindByID(id uint64) (*entities.LetterTemplate, error) {
	var template entities.LetterTemplate
	err := r.db.Preload("MediaAsset").First(&template, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &template, err
}

func (r *letterTemplateRepository) Update(template *entities.LetterTemplate) error {
	return r.db.Save(template).Error
}

func (r *letterTemplateRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.LetterTemplate{}, id).Error
}

func (r *letterTemplateRepository) FindMediaAsset(id uint64) (*entities.MediaAsset, error) {
	var media entities.MediaAsset
	err := r.db.First(&media, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &media, err
}

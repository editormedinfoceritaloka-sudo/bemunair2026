package repository

import (
	"errors"

	"bemunair2026/server/database/entities"
	"gorm.io/gorm"
)

type ListResult[T any] struct {
	Items      []T
	Total      int64
	Page       int
	PerPage    int
	TotalPages int
}

type Repository interface {
	ActiveCabinet() (*entities.CabinetTerm, error)
	CabinetBySlug(slug string, public bool) (*entities.CabinetTerm, error)
	Cabinets(page, perPage int) (ListResult[entities.CabinetTerm], error)
	CreateCabinet(value *entities.CabinetTerm) error
	UpdateCabinet(value *entities.CabinetTerm) error
	Units(cabinetID uint64, unitType string, public bool) ([]entities.Ministry, error)
	UnitBySlug(slug string, public bool) (*entities.Ministry, error)
	CreateUnit(value *entities.Ministry) error
	UpdateUnit(value *entities.Ministry) error
	Members(unitID uint64, public bool) ([]entities.OrganizationMember, error)
	CreateMember(value *entities.OrganizationMember) error
	UpdateMember(value *entities.OrganizationMember) error
	Programs(unitID uint64, page, perPage int, public bool) (ListResult[entities.WorkProgram], error)
	ProgramBySlug(unitSlug, programSlug string, public bool) (*entities.WorkProgram, error)
	CreateProgram(value *entities.WorkProgram) error
	UpdateProgram(value *entities.WorkProgram) error
	CreateMilestone(value *entities.WorkProgramMilestone) error
	CreateDocumentation(value *entities.WorkProgramDocumentation) error
	ReorderDocumentations(programID uint64, ids []uint64) error
	CreateMedia(value *entities.MediaAsset) error
	MediaByID(id uint64) (*entities.MediaAsset, error)
	DB() *gorm.DB
}

type cabinetRepository struct{ db *gorm.DB }

var _ Repository = (*cabinetRepository)(nil)

func New(db *gorm.DB) Repository { return &cabinetRepository{db: db} }

func (r *cabinetRepository) DB() *gorm.DB { return r.db }

func (r *cabinetRepository) ActiveCabinet() (*entities.CabinetTerm, error) {
	return r.findCabinet(r.db.Where("is_active = ? AND is_published = ?", true, true))
}

func (r *cabinetRepository) CabinetBySlug(slug string, public bool) (*entities.CabinetTerm, error) {
	query := r.db.Where("slug = ?", slug)
	if public {
		query = query.Where("is_active = ? AND is_published = ?", true, true)
	}
	return r.findCabinet(query)
}

func (r *cabinetRepository) findCabinet(query *gorm.DB) (*entities.CabinetTerm, error) {
	var value entities.CabinetTerm
	err := query.Preload("LogoMedia").Preload("HeroMedia").First(&value).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &value, err
}

func (r *cabinetRepository) Cabinets(page, perPage int) (ListResult[entities.CabinetTerm], error) {
	var values []entities.CabinetTerm
	var total int64
	query := r.db.Model(&entities.CabinetTerm{})
	if err := query.Count(&total).Error; err != nil {
		return ListResult[entities.CabinetTerm]{}, err
	}
	if err := query.Preload("LogoMedia").Order("created_at DESC").Limit(perPage).Offset((page - 1) * perPage).Find(&values).Error; err != nil {
		return ListResult[entities.CabinetTerm]{}, err
	}
	return ListResult[entities.CabinetTerm]{Items: values, Total: total, Page: page, PerPage: perPage, TotalPages: totalPages(total, perPage)}, nil
}

func (r *cabinetRepository) CreateCabinet(value *entities.CabinetTerm) error {
	return r.db.Create(value).Error
}
func (r *cabinetRepository) UpdateCabinet(value *entities.CabinetTerm) error {
	return r.db.Save(value).Error
}

func (r *cabinetRepository) Units(cabinetID uint64, unitType string, public bool) ([]entities.Ministry, error) {
	query := r.db.Where("cabinet_term_id = ?", cabinetID)
	if unitType != "" {
		query = query.Where("unit_type = ?", unitType)
	}
	if public {
		query = query.Where("is_active = ? AND is_published = ?", true, true)
	}
	var values []entities.Ministry
	err := query.Preload("LogoMedia").Preload("CoverMedia").Preload("Members", func(db *gorm.DB) *gorm.DB {
		q := db.Order("display_order ASC, id ASC")
		if public {
			q = q.Where("is_active = ?", true)
		}
		return q
	}).Preload("Members.PhotoMedia").Preload("Children", func(db *gorm.DB) *gorm.DB {
		q := db.Order("display_order ASC, id ASC")
		if public {
			q = q.Where("is_active = ? AND is_published = ?", true, true)
		}
		return q
	}).Preload("Children.LogoMedia").Order("display_order ASC, id ASC").Find(&values).Error
	return values, err
}

func (r *cabinetRepository) UnitBySlug(slug string, public bool) (*entities.Ministry, error) {
	query := r.db.Where("ministries.slug = ?", slug)
	if public {
		query = query.Joins("JOIN cabinet_terms ON cabinet_terms.id = ministries.cabinet_term_id").Where("ministries.is_active = ? AND ministries.is_published = ? AND cabinet_terms.is_active = ? AND cabinet_terms.is_published = ?", true, true, true, true)
	}
	var value entities.Ministry
	err := query.Preload("LogoMedia").Preload("CoverMedia").Preload("Members", func(db *gorm.DB) *gorm.DB {
		q := db.Order("display_order ASC, id ASC")
		if public {
			q = q.Where("is_active = ?", true)
		}
		return q
	}).Preload("Members.PhotoMedia").Preload("Children", func(db *gorm.DB) *gorm.DB {
		q := db.Order("display_order ASC, id ASC")
		if public {
			q = q.Where("is_active = ? AND is_published = ?", true, true)
		}
		return q
	}).Preload("Children.LogoMedia").First(&value).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &value, err
}

func (r *cabinetRepository) CreateUnit(value *entities.Ministry) error {
	return r.db.Create(value).Error
}
func (r *cabinetRepository) UpdateUnit(value *entities.Ministry) error { return r.db.Save(value).Error }

func (r *cabinetRepository) Members(unitID uint64, public bool) ([]entities.OrganizationMember, error) {
	query := r.db.Where("ministry_id = ?", unitID)
	if public {
		query = query.Where("is_active = ?", true)
	}
	var values []entities.OrganizationMember
	err := query.Preload("PhotoMedia").Order("display_order ASC, id ASC").Find(&values).Error
	return values, err
}

func (r *cabinetRepository) CreateMember(value *entities.OrganizationMember) error {
	return r.db.Create(value).Error
}
func (r *cabinetRepository) UpdateMember(value *entities.OrganizationMember) error {
	return r.db.Save(value).Error
}

func (r *cabinetRepository) Programs(unitID uint64, page, perPage int, public bool) (ListResult[entities.WorkProgram], error) {
	query := r.db.Model(&entities.WorkProgram{}).Where("ministry_id = ?", unitID)
	if public {
		query = query.Where("is_published = ?", true)
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return ListResult[entities.WorkProgram]{}, err
	}
	var values []entities.WorkProgram
	err := query.Preload("Ministry").Preload("CoverMedia").Order("display_order ASC, id ASC").Limit(perPage).Offset((page - 1) * perPage).Find(&values).Error
	if err != nil {
		return ListResult[entities.WorkProgram]{}, err
	}
	return ListResult[entities.WorkProgram]{Items: values, Total: total, Page: page, PerPage: perPage, TotalPages: totalPages(total, perPage)}, nil
}

func (r *cabinetRepository) ProgramBySlug(unitSlug, programSlug string, public bool) (*entities.WorkProgram, error) {
	query := r.db.Joins("JOIN ministries ON ministries.id = work_programs.ministry_id").Joins("JOIN cabinet_terms ON cabinet_terms.id = ministries.cabinet_term_id").Where("ministries.slug = ? AND work_programs.slug = ?", unitSlug, programSlug)
	if public {
		query = query.Where("ministries.is_active = ? AND ministries.is_published = ? AND cabinet_terms.is_active = ? AND cabinet_terms.is_published = ? AND work_programs.is_published = ?", true, true, true, true, true)
	}
	var value entities.WorkProgram
	err := query.Preload("Ministry").Preload("CoverMedia").Preload("Milestones", func(db *gorm.DB) *gorm.DB { return db.Order("display_order ASC, id ASC") }).Preload("Documentations", func(db *gorm.DB) *gorm.DB { return db.Order("display_order ASC, id ASC") }).Preload("Documentations.MediaAsset").First(&value).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &value, err
}

func (r *cabinetRepository) CreateProgram(value *entities.WorkProgram) error {
	return r.db.Create(value).Error
}
func (r *cabinetRepository) UpdateProgram(value *entities.WorkProgram) error {
	return r.db.Save(value).Error
}
func (r *cabinetRepository) CreateMilestone(value *entities.WorkProgramMilestone) error {
	return r.db.Create(value).Error
}

func (r *cabinetRepository) CreateDocumentation(value *entities.WorkProgramDocumentation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if value.IsCover {
			tx.Model(&entities.WorkProgramDocumentation{}).Where("work_program_id = ?", value.WorkProgramID).Update("is_cover", false)
		}
		return tx.Create(value).Error
	})
}

func (r *cabinetRepository) ReorderDocumentations(programID uint64, ids []uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for index, id := range ids {
			if err := tx.Model(&entities.WorkProgramDocumentation{}).Where("id = ? AND work_program_id = ?", id, programID).Update("display_order", index).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *cabinetRepository) CreateMedia(value *entities.MediaAsset) error {
	return r.db.Create(value).Error
}

func (r *cabinetRepository) MediaByID(id uint64) (*entities.MediaAsset, error) {
	var value entities.MediaAsset
	err := r.db.First(&value, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &value, err
}

func totalPages(total int64, perPage int) int {
	if total == 0 || perPage == 0 {
		return 0
	}
	return int((total + int64(perPage) - 1) / int64(perPage))
}

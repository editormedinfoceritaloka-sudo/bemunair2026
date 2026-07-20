package repository

import (
	"errors"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/pkg/constants"
	"gorm.io/gorm"
)

type ListParams struct {
	Page    int
	PerPage int

	PublishedOnly bool
}

type ArticleRepository interface {
	Create(article *entities.Article) error
	List(params ListParams) ([]entities.Article, int64, error)
	FindByID(id uint64) (*entities.Article, error)
	FindPublishedBySlug(slug string) (*entities.Article, error)
	ExistsBySlug(slug string, excludeID uint64) (bool, error)
	Update(article *entities.Article) error
	Delete(id uint64) error
}

type articleRepository struct {
	db *gorm.DB
}

var _ ArticleRepository = (*articleRepository)(nil)

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(article *entities.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) List(params ListParams) ([]entities.Article, int64, error) {
	var rows []entities.Article
	var total int64

	query := r.db.Model(&entities.Article{})
	if params.PublishedOnly {
		query = query.Where("status = ?", constants.ArticleStatusPublished)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	order := "created_at DESC"
	if params.PublishedOnly {
		order = "published_at DESC"
	}

	offset := (params.Page - 1) * params.PerPage
	err := query.
		Preload("Author").
		Order(order).
		Limit(params.PerPage).
		Offset(offset).
		Find(&rows).Error
	return rows, total, err
}

func (r *articleRepository) FindByID(id uint64) (*entities.Article, error) {
	var article entities.Article
	err := r.db.Preload("Author").First(&article, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &article, err
}

func (r *articleRepository) FindPublishedBySlug(slug string) (*entities.Article, error) {
	var article entities.Article
	err := r.db.Preload("Author").
		Where("slug = ? AND status = ?", slug, constants.ArticleStatusPublished).
		First(&article).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &article, err
}

func (r *articleRepository) ExistsBySlug(slug string, excludeID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&entities.Article{}).
		Where("slug = ? AND id <> ?", slug, excludeID).
		Count(&count).Error
	return count > 0, err
}

func (r *articleRepository) Update(article *entities.Article) error {
	return r.db.Save(article).Error
}

func (r *articleRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.Article{}, id).Error
}

package service

import (
	"errors"
	"fmt"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/article/dto"
	"bemunair2026/server/modules/article/repository"
	"bemunair2026/server/pkg/constants"
	"bemunair2026/server/pkg/utils"
)

const (
	defaultPage    = 1
	defaultPerPage = 10
	maxPerPage     = 50
)

type ListResult struct {
	Items      []dto.ArticleListItem
	Page       int
	PerPage    int
	Total      int64
	TotalPages int
}

type ArticleService interface {
	Create(req dto.CreateRequest, authorID uint64) (*dto.ArticleResponse, error)
	ListPublished(page, perPage int) (*ListResult, error)
	ListAll(page, perPage int) (*ListResult, error)
	GetPublishedBySlug(slug string) (*dto.ArticleResponse, error)
	GetByID(id uint64) (*dto.ArticleResponse, error)
	Update(id uint64, req dto.UpdateRequest) (*dto.ArticleResponse, error)
	SetPublished(id uint64, published bool) (*dto.ArticleResponse, error)
	Delete(id uint64) error
}

type articleService struct {
	repository repository.ArticleRepository
}

var _ ArticleService = (*articleService)(nil)

func NewArticleService(repository repository.ArticleRepository) ArticleService {
	return &articleService{repository: repository}
}

func (s *articleService) Create(req dto.CreateRequest, authorID uint64) (*dto.ArticleResponse, error) {
	if req.Title == "" {
		return nil, errors.New("title wajib diisi")
	}
	if req.Body == "" {
		return nil, errors.New("body wajib diisi")
	}

	slug, err := s.uniqueSlug(req.Title, 0)
	if err != nil {
		return nil, err
	}

	article := &entities.Article{
		Slug:       slug,
		Title:      req.Title,
		Excerpt:    req.Excerpt,
		Body:       utils.SanitizeArticleHTML(req.Body),
		CoverImage: req.CoverImage,
		AuthorID:   authorID,
		Status:     constants.ArticleStatusDraft,
	}
	if err := s.repository.Create(article); err != nil {
		return nil, err
	}
	res := dto.NewArticleResponse(article)
	return &res, nil
}

func (s *articleService) ListPublished(page, perPage int) (*ListResult, error) {
	return s.list(page, perPage, true)
}

func (s *articleService) ListAll(page, perPage int) (*ListResult, error) {
	return s.list(page, perPage, false)
}

func (s *articleService) list(page, perPage int, publishedOnly bool) (*ListResult, error) {
	page, perPage = normalizePaging(page, perPage)
	rows, total, err := s.repository.List(repository.ListParams{
		Page:          page,
		PerPage:       perPage,
		PublishedOnly: publishedOnly,
	})
	if err != nil {
		return nil, err
	}
	return &ListResult{
		Items:      dto.NewArticleListItems(rows),
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages(total, perPage),
	}, nil
}

func (s *articleService) GetPublishedBySlug(slug string) (*dto.ArticleResponse, error) {
	article, err := s.repository.FindPublishedBySlug(slug)
	if err != nil || article == nil {
		return nil, err
	}
	res := dto.NewArticleResponse(article)
	return &res, nil
}

func (s *articleService) GetByID(id uint64) (*dto.ArticleResponse, error) {
	article, err := s.repository.FindByID(id)
	if err != nil || article == nil {
		return nil, err
	}
	res := dto.NewArticleResponse(article)
	return &res, nil
}

func (s *articleService) Update(id uint64, req dto.UpdateRequest) (*dto.ArticleResponse, error) {
	article, err := s.repository.FindByID(id)
	if err != nil || article == nil {
		return nil, err
	}
	if req.Title == "" {
		return nil, errors.New("title wajib diisi")
	}
	if req.Body == "" {
		return nil, errors.New("body wajib diisi")
	}

	if req.Title != article.Title {
		slug, err := s.uniqueSlug(req.Title, article.ID)
		if err != nil {
			return nil, err
		}
		article.Slug = slug
	}
	article.Title = req.Title
	article.Excerpt = req.Excerpt
	article.Body = utils.SanitizeArticleHTML(req.Body)
	article.CoverImage = req.CoverImage

	if err := s.repository.Update(article); err != nil {
		return nil, err
	}
	res := dto.NewArticleResponse(article)
	return &res, nil
}

func (s *articleService) SetPublished(id uint64, published bool) (*dto.ArticleResponse, error) {
	article, err := s.repository.FindByID(id)
	if err != nil || article == nil {
		return nil, err
	}

	if published {
		article.Status = constants.ArticleStatusPublished
		if article.PublishedAt == nil {
			now := time.Now()
			article.PublishedAt = &now
		}
	} else {
		article.Status = constants.ArticleStatusDraft
		article.PublishedAt = nil
	}

	if err := s.repository.Update(article); err != nil {
		return nil, err
	}
	res := dto.NewArticleResponse(article)
	return &res, nil
}

func (s *articleService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func (s *articleService) uniqueSlug(title string, excludeID uint64) (string, error) {
	base := utils.Slugify(title)
	if base == "" {
		base = "artikel"
	}
	slug := base
	for i := 2; ; i++ {
		exists, err := s.repository.ExistsBySlug(slug, excludeID)
		if err != nil {
			return "", err
		}
		if !exists {
			return slug, nil
		}
		slug = fmt.Sprintf("%s-%d", base, i)
	}
}

func normalizePaging(page, perPage int) (int, int) {
	if page < 1 {
		page = defaultPage
	}
	if perPage < 1 {
		perPage = defaultPerPage
	}
	if perPage > maxPerPage {
		perPage = maxPerPage
	}
	return page, perPage
}

func totalPages(total int64, perPage int) int {
	if perPage <= 0 || total <= 0 {
		return 0
	}
	return int((total + int64(perPage) - 1) / int64(perPage))
}

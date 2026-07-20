package dto

import (
	"time"

	"bemunair2026/server/database/entities"
)

type CreateRequest struct {
	Title      string  `json:"title"`
	Excerpt    *string `json:"excerpt"`
	Body       string  `json:"body"`
	CoverImage *string `json:"cover_image"`
}

type UpdateRequest struct {
	Title      string  `json:"title"`
	Excerpt    *string `json:"excerpt"`
	Body       string  `json:"body"`
	CoverImage *string `json:"cover_image"`
}

type PublishRequest struct {
	Published bool `json:"published"`
}

type AuthorSummary struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ArticleResponse struct {
	ID          uint64         `json:"id"`
	Slug        string         `json:"slug"`
	Title       string         `json:"title"`
	Excerpt     *string        `json:"excerpt,omitempty"`
	Body        string         `json:"body"`
	CoverImage  *string        `json:"cover_image,omitempty"`
	AuthorID    uint64         `json:"author_id"`
	Author      *AuthorSummary `json:"author,omitempty"`
	Status      string         `json:"status"`
	PublishedAt *time.Time     `json:"published_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ArticleListItem struct {
	ID          uint64         `json:"id"`
	Slug        string         `json:"slug"`
	Title       string         `json:"title"`
	Excerpt     *string        `json:"excerpt,omitempty"`
	CoverImage  *string        `json:"cover_image,omitempty"`
	AuthorID    uint64         `json:"author_id"`
	Author      *AuthorSummary `json:"author,omitempty"`
	Status      string         `json:"status"`
	PublishedAt *time.Time     `json:"published_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func NewArticleResponse(article *entities.Article) ArticleResponse {
	if article == nil {
		return ArticleResponse{}
	}

	return ArticleResponse{
		ID:          article.ID,
		Slug:        article.Slug,
		Title:       article.Title,
		Excerpt:     article.Excerpt,
		Body:        article.Body,
		CoverImage:  article.CoverImage,
		AuthorID:    article.AuthorID,
		Author:      newAuthorSummary(article.Author),
		Status:      article.Status,
		PublishedAt: article.PublishedAt,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
}

func NewArticleListItems(articles []entities.Article) []ArticleListItem {
	items := make([]ArticleListItem, 0, len(articles))
	for i := range articles {
		a := &articles[i]
		items = append(items, ArticleListItem{
			ID:          a.ID,
			Slug:        a.Slug,
			Title:       a.Title,
			Excerpt:     a.Excerpt,
			CoverImage:  a.CoverImage,
			AuthorID:    a.AuthorID,
			Author:      newAuthorSummary(a.Author),
			Status:      a.Status,
			PublishedAt: a.PublishedAt,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
		})
	}
	return items
}

func newAuthorSummary(user *entities.User) *AuthorSummary {
	if user == nil {
		return nil
	}
	return &AuthorSummary{ID: user.ID, Name: user.Name}
}

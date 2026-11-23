package submission_support

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
)

func TestArticleHandoffCreatesAndReopensDraft(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = sqlDB.Close() })
	sqlDB.SetMaxOpenConns(1)
	for _, statement := range []string{
		`CREATE TABLE content_submissions (id INTEGER PRIMARY KEY, title TEXT, service_type TEXT, status TEXT, article_drive_link TEXT)`,
		`CREATE TABLE articles (id INTEGER PRIMARY KEY, slug TEXT UNIQUE, title TEXT, excerpt TEXT, body TEXT, cover_image TEXT, author_id INTEGER, source_submission_id INTEGER UNIQUE, status TEXT, published_at DATETIME, created_at DATETIME, updated_at DATETIME)`,
		`INSERT INTO content_submissions VALUES (1, 'Kabar Kementerian', 'ARTICLE', 'APPROVED', 'https://docs.google.com/document/d/contoh')`,
	} {
		if err := db.Exec(statement).Error; err != nil {
			t.Fatal(err)
		}
	}

	router := gin.New()
	handler := &controller{db: db}
	router.POST("/submissions/:id/draft", func(ctx *gin.Context) {
		ctx.Set("claims", &middlewares.Claims{UserID: 42})
		handler.createArticleDraft(ctx)
	})

	for _, status := range []int{http.StatusCreated, http.StatusOK} {
		response := httptest.NewRecorder()
		router.ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/submissions/1/draft", nil))
		var payload struct {
			Data struct {
				ID      uint64 `json:"id"`
				Created bool   `json:"created"`
			} `json:"data"`
		}
		if response.Code != status {
			t.Fatalf("status = %d, want %d: %s", response.Code, status, response.Body.String())
		}
		if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
			t.Fatal(err)
		}
		if payload.Data.ID != 1 || payload.Data.Created != (status == http.StatusCreated) {
			t.Fatalf("unexpected response: %s", response.Body.String())
		}
	}

	var draft entities.Article
	if err := db.First(&draft).Error; err != nil {
		t.Fatal(err)
	}
	if draft.Status != "DRAFT" || draft.AuthorID != 42 || draft.SourceSubmissionID == nil || *draft.SourceSubmissionID != 1 {
		t.Fatalf("unexpected draft metadata: %+v", draft)
	}
	if strings.Contains(draft.Body, "Kepala berita") || !strings.Contains(draft.Body, "https://docs.google.com/document/d/contoh") {
		t.Fatalf("draft should contain the source link without the template: %s", draft.Body)
	}
}

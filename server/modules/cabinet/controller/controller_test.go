package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/cabinet"
	"bemunair2026/server/modules/cabinet/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T, name string) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file:"+name+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&entities.CabinetTerm{}, &entities.Ministry{}, &entities.WorkProgram{}, &entities.WorkProgramMilestone{}, &entities.WorkProgramDocumentation{}, &entities.MediaAsset{}); err != nil {
		t.Fatal(err)
	}
	fixTimestampColumns(t, db)
	return db
}

// fixTimestampColumns mengubah kolom bertipe "timestamp with time zone" (tag
// entity untuk database produksi) menjadi "datetime" agar driver SQLite dapat
// meng-scan nilainya sebagai time.Time. Hanya untuk test.
func fixTimestampColumns(t *testing.T, db *gorm.DB) {
	t.Helper()
	tables := []string{"cabinet_terms", "ministries", "work_programs", "work_program_milestones", "work_program_documentations", "media_assets"}
	for _, table := range tables {
		for _, col := range []string{"created_at", "updated_at"} {
			sqls := []string{
				fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s_tmp", table, col, col),
				fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s datetime", table, col),
				fmt.Sprintf("UPDATE %s SET %s = %s_tmp", table, col, col),
				fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s_tmp", table, col),
			}
			for _, s := range sqls {
				if err := db.Exec(s).Error; err != nil {
					t.Fatalf("fix %s.%s: %v", table, col, err)
				}
			}
		}
	}
}

func seedPublishedProgram(t *testing.T, db *gorm.DB) {
	t.Helper()
	cabinetTerm := entities.CabinetTerm{Name: "Kabinet Cerita Loka", Slug: "cerita-loka", IsActive: true, IsPublished: true}
	if err := db.Create(&cabinetTerm).Error; err != nil {
		t.Fatal(err)
	}
	ministry := entities.Ministry{Code: "KEM1", Name: "Kementerian Satu", Slug: "kementerian-satu", CabinetTermID: &cabinetTerm.ID, UnitType: "KEMENTERIAN", IsActive: true, IsPublished: true}
	if err := db.Create(&ministry).Error; err != nil {
		t.Fatal(err)
	}
	program := entities.WorkProgram{MinistryID: ministry.ID, Name: "Cerita Visual", Slug: "cerita-visual", LifecycleStatus: "ONGOING", IsPublished: true}
	if err := db.Create(&program).Error; err != nil {
		t.Fatal(err)
	}
}

func newTestRouter(t *testing.T, name string) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	db := newTestDB(t, name)
	seedPublishedProgram(t, db)
	router := gin.New()
	api := router.Group("/api/v1")
	cabinet.RegisterRoutes(api, repository.New(db), "test-secret")
	return router
}

func doRequest(router *gin.Engine, path string) (*httptest.ResponseRecorder, map[string]any) {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	var body map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &body)
	return rec, body
}

func TestPublicProgramBySlugRoute(t *testing.T) {
	router := newTestRouter(t, "ctl_program_by_slug")

	rec, body := doRequest(router, "/api/v1/cabinet/programs/cerita-visual")
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rec.Code, rec.Body.String())
	}
	data, ok := body["data"].(map[string]any)
	if !ok {
		t.Fatalf("expected data object, got %v", body)
	}
	if data["slug"] != "cerita-visual" || data["name"] != "Cerita Visual" {
		t.Fatalf("unexpected data: %v", data)
	}
}

func TestPublicProgramBySlugRouteNotFound(t *testing.T) {
	router := newTestRouter(t, "ctl_program_by_slug_notfound")

	rec, _ := doRequest(router, "/api/v1/cabinet/programs/tidak-ada")
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestPublicProgramOldRouteStillWorks(t *testing.T) {
	router := newTestRouter(t, "ctl_program_old_route")

	rec, body := doRequest(router, "/api/v1/cabinet/units/kementerian-satu/programs/cerita-visual")
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "Cerita Visual") {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
	_ = body
}

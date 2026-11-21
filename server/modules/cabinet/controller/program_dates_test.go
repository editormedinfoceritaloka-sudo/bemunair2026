package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/cabinet"
	"bemunair2026/server/modules/cabinet/repository"
	"bemunair2026/server/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func newAdminRouter(t *testing.T, name string) (*gin.Engine, *gorm.DB) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	db := newTestDB(t, name)
	secret := "test-secret"
	router := gin.New()
	cabinet.RegisterRoutes(router.Group("/api/v1"), repository.New(db), secret)
	return router, db
}

func adminRequest(t *testing.T, router *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(1),
		"role":    constants.RoleAdminMedinfo,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString([]byte("test-secret"))
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+signed)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestCreateProgramStoresDates(t *testing.T) {
	router, db := newAdminRouter(t, "program_dates")
	seedPublishedProgram(t, db)
	var ministry entities.Ministry
	if err := db.First(&ministry).Error; err != nil {
		t.Fatal(err)
	}

	body := `{"ministry_id":` + strconv.FormatUint(ministry.ID, 10) + `,"name":"Program Bertanggal","status":"PLANNED","start_date":"2026-03-01","end_date":"2026-04-30"}`
	w := adminRequest(t, router, http.MethodPost, "/api/v1/admin/work-programs", body)
	if w.Code != http.StatusCreated {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}

	var program entities.WorkProgram
	if err := db.Where("name = ?", "Program Bertanggal").First(&program).Error; err != nil {
		t.Fatal(err)
	}
	if program.StartDate == nil || program.StartDate.Format("2006-01-02") != "2026-03-01" {
		t.Fatalf("start_date=%v", program.StartDate)
	}
	if program.EndDate == nil || program.EndDate.Format("2006-01-02") != "2026-04-30" {
		t.Fatalf("end_date=%v", program.EndDate)
	}
}

func TestCreateProgramRejectsReversedDates(t *testing.T) {
	router, db := newAdminRouter(t, "program_dates_invalid")
	seedPublishedProgram(t, db)
	var ministry entities.Ministry
	if err := db.First(&ministry).Error; err != nil {
		t.Fatal(err)
	}

	body := `{"ministry_id":` + strconv.FormatUint(ministry.ID, 10) + `,"name":"Rentang Terbalik","status":"PLANNED","start_date":"2026-05-01","end_date":"2026-04-01"}`
	w := adminRequest(t, router, http.MethodPost, "/api/v1/admin/work-programs", body)
	if w.Code == http.StatusCreated {
		t.Fatalf("expected rejection, got status=%d body=%s", w.Code, w.Body.String())
	}
	var count int64
	if err := db.Model(&entities.WorkProgram{}).Where("name = ?", "Rentang Terbalik").Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("invalid program was stored: count=%d", count)
	}
}

func TestCreateProgramGeneratesUniqueSlug(t *testing.T) {
	router, db := newAdminRouter(t, "program_slug")
	seedPublishedProgram(t, db)
	var ministry entities.Ministry
	if err := db.First(&ministry).Error; err != nil {
		t.Fatal(err)
	}

	body := `{"ministry_id":` + strconv.FormatUint(ministry.ID, 10) + `,"name":"Program Tanpa Slug","status":"PLANNED"}`
	if w := adminRequest(t, router, http.MethodPost, "/api/v1/admin/work-programs", body); w.Code != http.StatusCreated {
		t.Fatalf("first create status=%d body=%s", w.Code, w.Body.String())
	}
	var first entities.WorkProgram
	if err := db.Where("name = ?", "Program Tanpa Slug").First(&first).Error; err != nil {
		t.Fatal(err)
	}
	if first.Slug != "program-tanpa-slug" {
		t.Fatalf("slug=%q", first.Slug)
	}

	if w := adminRequest(t, router, http.MethodPost, "/api/v1/admin/work-programs", body); w.Code != http.StatusCreated {
		t.Fatalf("second create status=%d body=%s", w.Code, w.Body.String())
	}
	var second entities.WorkProgram
	if err := db.Where("name = ? AND id <> ?", "Program Tanpa Slug", first.ID).First(&second).Error; err != nil {
		t.Fatal(err)
	}
	if second.Slug != "program-tanpa-slug-1" {
		t.Fatalf("slug=%q", second.Slug)
	}
}

package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/cabinet"
	"bemunair2026/server/modules/cabinet/repository"
	"bemunair2026/server/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestDeleteProgramRemovesChildrenKeepsMedia(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := newTestDB(t, "delete_program")
	seedPublishedProgram(t, db)

	var program entities.WorkProgram
	if err := db.Where("slug = ?", "cerita-visual").First(&program).Error; err != nil {
		t.Fatal(err)
	}
	media := entities.MediaAsset{ImageKitFileID: "file-1", URL: "/uploads/dokumentasi.jpg", Name: "dokumentasi.jpg", AltText: "Dokumentasi", MimeType: "image/jpeg", Purpose: "work_program_documentation"}
	if err := db.Create(&media).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&entities.WorkProgramMilestone{WorkProgramID: program.ID, Title: "Milestone 1"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&entities.WorkProgramDocumentation{WorkProgramID: program.ID, MediaAssetID: media.ID}).Error; err != nil {
		t.Fatal(err)
	}

	secret := "test-secret"
	router := gin.New()
	cabinet.RegisterRoutes(router.Group("/api/v1"), repository.New(db), secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(1),
		"role":    constants.RoleAdminMedinfo,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/admin/work-programs/%d", program.ID), nil)
	req.Header.Set("Authorization", "Bearer "+signed)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", w.Code, w.Body.String())
	}
	for table, want := range map[string]int64{"work_programs": 0, "work_program_milestones": 0, "work_program_documentations": 0, "media_assets": 1} {
		var count int64
		if err := db.Table(table).Count(&count).Error; err != nil {
			t.Fatal(err)
		}
		if count != want {
			t.Fatalf("%s count=%d want=%d", table, count, want)
		}
	}
}

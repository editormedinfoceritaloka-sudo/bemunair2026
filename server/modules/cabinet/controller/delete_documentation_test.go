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

func TestDeleteDocumentationScopesDeletionToProgram(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := newTestDB(t, "delete_documentation")
	seedPublishedProgram(t, db)
	var program entities.WorkProgram
	if err := db.Where("slug = ?", "cerita-visual").First(&program).Error; err != nil {
		t.Fatal(err)
	}
	other := entities.WorkProgram{MinistryID: program.MinistryID, Name: "Program lain", Slug: "program-lain"}
	if err := db.Create(&other).Error; err != nil {
		t.Fatal(err)
	}
	media := entities.MediaAsset{ImageKitFileID: "shared-photo", URL: "/uploads/photo.jpg", Name: "photo.jpg", AltText: "Foto", MimeType: "image/jpeg"}
	if err := db.Create(&media).Error; err != nil {
		t.Fatal(err)
	}
	photos := []entities.WorkProgramDocumentation{
		{WorkProgramID: program.ID, MediaAssetID: media.ID},
		{WorkProgramID: other.ID, MediaAssetID: media.ID},
	}
	if err := db.Create(&photos).Error; err != nil {
		t.Fatal(err)
	}

	secret := "test-secret"
	router := gin.New()
	cabinet.RegisterRoutes(router.Group("/api/v1"), repository.New(db), secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1, "role": constants.RoleAdminMedinfo, "exp": time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatal(err)
	}

	for _, step := range []struct {
		programID uint64
		status    int
		remaining int64
	}{
		{other.ID, http.StatusNotFound, 2},
		{program.ID, http.StatusOK, 1},
		{program.ID, http.StatusNotFound, 1},
	} {
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/admin/work-programs/%d/documentations/%d", step.programID, photos[0].ID), nil)
		req.Header.Set("Authorization", "Bearer "+signed)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)
		if response.Code != step.status {
			t.Fatalf("status=%d want=%d body=%s", response.Code, step.status, response.Body.String())
		}
		var count int64
		if err := db.Model(&entities.WorkProgramDocumentation{}).Count(&count).Error; err != nil || count != step.remaining {
			t.Fatalf("documentation count=%d want=%d err=%v", count, step.remaining, err)
		}
	}
	var count int64
	if err := db.Model(&entities.MediaAsset{}).Where("id = ?", media.ID).Count(&count).Error; err != nil || count != 1 {
		t.Fatalf("shared media count=%d err=%v", count, err)
	}
}

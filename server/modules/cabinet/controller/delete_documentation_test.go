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
	otherMinistry := entities.Ministry{Code: "KEM2", Name: "Kementerian Dua", Slug: "kementerian-dua"}
	if err := db.Create(&otherMinistry).Error; err != nil {
		t.Fatal(err)
	}
	other := entities.WorkProgram{MinistryID: otherMinistry.ID, Name: "Program lain", Slug: "program-lain"}
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
	// The permission check reads only these columns. SQLite lacks MySQL's user role ENUM.
	for _, statement := range []string{
		"CREATE TABLE users (id INTEGER PRIMARY KEY, ministry_id INTEGER)",
		"CREATE TABLE user_organization_roles (user_id INTEGER, ministry_id INTEGER, permission TEXT, is_active BOOLEAN)",
	} {
		if err := db.Exec(statement).Error; err != nil {
			t.Fatal(err)
		}
	}
	const ministryUserID = 1
	if err := db.Exec("INSERT INTO users (id, ministry_id) VALUES (?, ?)", ministryUserID, program.MinistryID).Error; err != nil {
		t.Fatal(err)
	}

	secret := "test-secret"
	router := gin.New()
	cabinet.RegisterRoutes(router.Group("/api/v1"), repository.New(db), secret)

	for _, step := range []struct {
		name            string
		role            string
		programID       uint64
		documentationID uint64
		status          int
		remaining       int64
	}{
		{"ministry cannot delete from another ministry", constants.RoleAdmin, other.ID, photos[1].ID, http.StatusForbidden, 2},
		{"medinfo cannot delete with a mismatched program", constants.RoleAdminMedinfo, other.ID, photos[0].ID, http.StatusNotFound, 2},
		{"ministry can delete from its own program", constants.RoleAdmin, program.ID, photos[0].ID, http.StatusOK, 1},
		{"ministry cannot delete the same photo twice", constants.RoleAdmin, program.ID, photos[0].ID, http.StatusNotFound, 1},
		{"medinfo can delete across ministries", constants.RoleAdminMedinfo, other.ID, photos[1].ID, http.StatusOK, 0},
	} {
		t.Run(step.name, func(t *testing.T) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id": ministryUserID, "role": step.role, "exp": time.Now().Add(time.Hour).Unix(),
			})
			signed, err := token.SignedString([]byte(secret))
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/admin/work-programs/%d/documentations/%d", step.programID, step.documentationID), nil)
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
		})
	}
	var count int64
	if err := db.Model(&entities.MediaAsset{}).Where("id = ?", media.ID).Count(&count).Error; err != nil || count != 1 {
		t.Fatalf("shared media count=%d err=%v", count, err)
	}
}

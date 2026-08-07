package service

import (
	"fmt"
	"strings"
	"testing"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/cabinet/dto"
	"bemunair2026/server/modules/cabinet/repository"
	"bemunair2026/server/pkg/constants"
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
	cabinet := entities.CabinetTerm{Name: "Kabinet Cerita Loka", Slug: "cerita-loka", IsActive: true, IsPublished: true}
	if err := db.Create(&cabinet).Error; err != nil {
		t.Fatal(err)
	}
	ministry := entities.Ministry{Code: "KEM1", Name: "Kementerian Satu", Slug: "kementerian-satu", CabinetTermID: &cabinet.ID, UnitType: "KEMENTERIAN", IsActive: true, IsPublished: true}
	if err := db.Create(&ministry).Error; err != nil {
		t.Fatal(err)
	}
	program := entities.WorkProgram{MinistryID: ministry.ID, Name: "Cerita Visual", Slug: "cerita-visual", LifecycleStatus: "ONGOING", IsPublished: true}
	if err := db.Create(&program).Error; err != nil {
		t.Fatal(err)
	}
}

func testActor() *middlewares.Claims {
	return &middlewares.Claims{Role: constants.RoleAdminMedinfo}
}

func TestPublicProgramBySlugFound(t *testing.T) {
	db := newTestDB(t, "svc_found")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))

	value, err := svc.PublicProgramBySlug("cerita-visual")
	if err != nil || value == nil {
		t.Fatalf("expected program, got %v err=%v", value, err)
	}
	if value.Name != "Cerita Visual" || value.MinistryName != "Kementerian Satu" {
		t.Fatalf("unexpected program: %+v", value)
	}
}

func TestPublicProgramBySlugNotFound(t *testing.T) {
	db := newTestDB(t, "svc_notfound")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))

	value, err := svc.PublicProgramBySlug("tidak-ada")
	if err != nil || value != nil {
		t.Fatalf("expected nil,nil got %v err=%v", value, err)
	}
}

func TestCreateProgramAutoSlug(t *testing.T) {
	db := newTestDB(t, "svc_autoslug")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))

	value, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Cerita Visual", LifecycleStatus: "ONGOING"}, testActor())
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if !strings.HasPrefix(value.Slug, "cerita-visual-") {
		t.Fatalf("expected auto slug with datetime suffix, got %q", value.Slug)
	}
	if len(value.Slug) != len("cerita-visual-")+len("20060102-150405") {
		t.Fatalf("unexpected slug length: %q", value.Slug)
	}
	var count int64
	if err := db.Model(&entities.WorkProgram{}).Where("slug = ?", value.Slug).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("expected slug %q stored once, found %d", value.Slug, count)
	}
}

func TestCreateProgramUniqueSlugConflict(t *testing.T) {
	db := newTestDB(t, "svc_conflict")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))

	value, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Program Lain", Slug: "cerita-visual", LifecycleStatus: "ONGOING"}, testActor())
	if err == nil || !strings.Contains(err.Error(), "slug program sudah digunakan") {
		t.Fatalf("expected slug conflict error, got value=%v err=%v", value, err)
	}
}

func TestUpdateProgramPreserveSlug(t *testing.T) {
	db := newTestDB(t, "svc_preserve")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))

	value, err := svc.UpdateProgram(1, dto.ProgramRequest{MinistryID: 1, Name: "Cerita Visual", LifecycleStatus: "ONGOING"}, testActor())
	if err != nil || value == nil {
		t.Fatalf("update failed: %v", err)
	}
	if value.Slug != "cerita-visual" {
		t.Fatalf("expected slug preserved, got %q", value.Slug)
	}

	value, err = svc.UpdateProgram(1, dto.ProgramRequest{MinistryID: 1, Name: "Cerita Visual", Slug: "cerita-baru", LifecycleStatus: "ONGOING"}, testActor())
	if err != nil || value == nil {
		t.Fatalf("update with new slug failed: %v", err)
	}
	if value.Slug != "cerita-baru" {
		t.Fatalf("expected new slug applied, got %q", value.Slug)
	}
}

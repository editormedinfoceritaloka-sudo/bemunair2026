package repository

import (
	"fmt"
	"testing"

	"bemunair2026/server/database/entities"
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

func TestProgramBySlugOnlyFound(t *testing.T) {
	db := newTestDB(t, "repo_found")
	seedPublishedProgram(t, db)
	repo := New(db)

	value, err := repo.ProgramBySlugOnly("cerita-visual", true)
	if err != nil || value == nil {
		t.Fatalf("expected program, got %v err=%v", value, err)
	}
	if value.Name != "Cerita Visual" || value.Ministry == nil || value.Ministry.Name != "Kementerian Satu" {
		t.Fatalf("unexpected program: %+v", value)
	}
}

func TestProgramBySlugOnlyUnpublishedHidden(t *testing.T) {
	db := newTestDB(t, "repo_unpublished")
	seedPublishedProgram(t, db)
	if err := db.Model(&entities.WorkProgram{}).Where("slug = ?", "cerita-visual").Update("is_published", false).Error; err != nil {
		t.Fatal(err)
	}
	repo := New(db)

	value, err := repo.ProgramBySlugOnly("cerita-visual", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != nil {
		t.Fatalf("expected nil for unpublished program, got %+v", value)
	}
}

func TestProgramBySlugOnlyNotFound(t *testing.T) {
	db := newTestDB(t, "repo_notfound")
	seedPublishedProgram(t, db)
	repo := New(db)

	value, err := repo.ProgramBySlugOnly("tidak-ada", true)
	if err != nil || value != nil {
		t.Fatalf("expected nil,nil got %v err=%v", value, err)
	}
}

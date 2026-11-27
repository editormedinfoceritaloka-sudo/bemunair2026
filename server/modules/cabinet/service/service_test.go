package service

import (
	"errors"
	"fmt"
	"testing"
	"time"

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
	if value.Slug != "cerita-visual-1" {
		t.Fatalf("expected auto slug with numeric suffix, got %q", value.Slug)
	}
	var count int64
	if err := db.Model(&entities.WorkProgram{}).Where("slug = ?", value.Slug).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("expected slug %q stored once, found %d", value.Slug, count)
	}
}

func TestCreateProgramSuffixesDuplicateSlug(t *testing.T) {
	db := newTestDB(t, "svc_conflict")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))

	value, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Program Lain", Slug: "cerita-visual", LifecycleStatus: "ONGOING"}, testActor())
	if err != nil || value == nil {
		t.Fatalf("create failed: %v", err)
	}
	if value.Slug != "cerita-visual-1" {
		t.Fatalf("expected suffixed slug, got %q", value.Slug)
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

func TestSeederOrganizationTypesAccepted(t *testing.T) {
	for _, unitType := range []string{"MENKO", "BPII", "KEMENTERIAN"} {
		req := dto.UnitRequest{Code: "TEST", Name: "Unit Test", UnitType: unitType}
		if err := validateUnit(req); err != nil {
			t.Fatalf("expected seeder unit type %s to be accepted: %v", unitType, err)
		}
	}
}

func TestSeederPositionTypesAccepted(t *testing.T) {
	for _, positionType := range []string{"MENKO", "MENTERI", "DIRJEN", "KABIRO", "KEPALA", "DEPUTI"} {
		req := dto.MemberRequest{Name: "Anggota", Position: "Jabatan", PositionType: positionType}
		if err := validateMember(req); err != nil {
			t.Fatalf("expected seeder position type %s to be accepted: %v", positionType, err)
		}
	}
}

func TestDeleteUnitRequiresMedinfo(t *testing.T) {
	svc := New(repository.New(newTestDB(t, "svc_delete_access")))
	actor := &middlewares.Claims{Role: constants.RoleAdmin}

	deleted, err := svc.DeleteUnit(1, actor)
	if deleted || !errors.Is(err, ErrUnitAccessDenied) {
		t.Fatalf("expected access denied, got deleted=%v err=%v", deleted, err)
	}
}

func TestMemberToDTOIncludesActiveStatus(t *testing.T) {
	value := memberToDTO(entities.OrganizationMember{ID: 1, Name: "Anggota", IsActive: true})
	if !value.IsActive {
		t.Fatal("expected active status in member response")
	}
}

func TestProgramVisibilityFollowsStatus(t *testing.T) {
	db := newTestDB(t, "svc_visibility")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))

	draft, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Program Draft", LifecycleStatus: "DRAFT"}, testActor())
	if err != nil {
		t.Fatal(err)
	}
	if draft.IsPublished {
		t.Fatal("draft should not be published")
	}

	planned, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Program Rencana", LifecycleStatus: "PLANNED"}, testActor())
	if err != nil {
		t.Fatal(err)
	}
	if !planned.IsPublished {
		t.Fatal("planned should be published")
	}

	updated, err := svc.UpdateProgram(planned.ID, dto.ProgramRequest{MinistryID: 1, Name: "Program Rencana", LifecycleStatus: "DRAFT"}, testActor())
	if err != nil {
		t.Fatal(err)
	}
	if updated.IsPublished {
		t.Fatal("draft update should unpublish")
	}
}

func TestProgramStatusDerivedFromDates(t *testing.T) {
	db := newTestDB(t, "svc_derived")
	seedPublishedProgram(t, db)
	svc := New(repository.New(db))
	now := time.Now()

	futureStart := now.AddDate(0, 0, 7).Format("2006-01-02")
	futureEnd := now.AddDate(0, 0, 14).Format("2006-01-02")
	planned, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Akan Datang", LifecycleStatus: "PLANNED", StartDate: &futureStart, EndDate: &futureEnd}, testActor())
	if err != nil {
		t.Fatal(err)
	}
	if planned.Status != "PLANNED" {
		t.Fatalf("status=%q", planned.Status)
	}

	pastStart := now.AddDate(0, 0, -2).Format("2006-01-02")
	nearEnd := now.AddDate(0, 0, 2).Format("2006-01-02")
	ongoing, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Sedang Jalan", LifecycleStatus: "ONGOING", StartDate: &pastStart, EndDate: &nearEnd}, testActor())
	if err != nil {
		t.Fatal(err)
	}
	if ongoing.Status != "ONGOING" {
		t.Fatalf("status=%q", ongoing.Status)
	}

	oldStart := now.AddDate(0, 0, -14).Format("2006-01-02")
	oldEnd := now.AddDate(0, 0, -7).Format("2006-01-02")
	completed, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: 1, Name: "Sudah Selesai", LifecycleStatus: "COMPLETED", StartDate: &oldStart, EndDate: &oldEnd}, testActor())
	if err != nil {
		t.Fatal(err)
	}
	if completed.Status != "COMPLETED" {
		t.Fatalf("status=%q", completed.Status)
	}
}

func TestMemberPortraitSource(t *testing.T) {
	db := newTestDB(t, "member_portrait_source")
	s := cabinetService{repo: repository.New(db)}
	if err := s.validatePortrait(nil); err != nil {
		t.Fatal(err)
	}
	for i, tc := range []struct {
		url   string
		valid bool
	}{
		{"https://res.cloudinary.com/cloud/image/upload/v1/person.jpg", true},
		{"/uploads/profiles/person.jpg", false},
		{"https://res.cloudinary.com.example/cloud/image/upload/person.jpg", false},
		{"https://res.cloudinary.com/cloud/video/upload/person.mp4", false},
	} {
		photo := entities.MediaAsset{ImageKitFileID: fmt.Sprint(i), URL: tc.url, Name: "portrait", AltText: "portrait", MimeType: "image/jpeg", Purpose: "ORGANIZATION_MEMBER"}
		if err := db.Create(&photo).Error; err != nil {
			t.Fatal(err)
		}
		if err := s.validatePortrait(&photo.ID); (err == nil) != tc.valid {
			t.Errorf("validatePortrait(%q) = %v", tc.url, err)
		}
	}
}

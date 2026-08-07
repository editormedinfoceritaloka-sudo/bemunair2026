# Find Program Kerja by Program Slug — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Halaman publik `/program-kerja/[slug]` dapat mengambil detail program kerja cukup dengan slug program saja, lewat endpoint publik baru `GET /api/v1/cabinet/programs/:programSlug`.

**Architecture:** Tambah method repository `ProgramBySlugOnly` (join ministry + cabinet term, filter published), service `PublicProgramBySlug`, handler + route baru di group publik `/cabinet`. Slug otomatis saat create diberi suffix datetime (`YYYYMMDD-HHMMSS`) agar unik global; update mempertahankan slug lama; slug manual divalidasi unik. Client memanggil endpoint baru. Endpoint lama dua-slug tetap dipertahankan.

**Tech Stack:** Go 1.25 + Gin + GORM (sqlite driver untuk test), SvelteKit + TypeScript + pnpm, dokumentasi `docs/api/`.

## Global Constraints

- Endpoint lama `GET /api/v1/cabinet/units/:slug/programs/:programSlug` TIDAK dihapus (tetap terdaftar dan berfungsi).
- `WorkProgramCard.svelte` (komponen tidak terpakai) KELUAR dari scope — jangan disentuh.
- Tidak ada migration/unique index DB — validasi keunikan cukup di service.
- Format datetime slug: `time.Now().Format("20060102-150405")` (waktu lokal server, contoh: `cerita-visual-20260807-153045`).
- Slug manual yang dikirim admin dipakai apa adanya (tanpa suffix datetime).
- Setiap task ditutup dengan validasi: Go → `gofmt` file berubah + `go test ./... -v`; client → `cd client && pnpm check`.
- Test memakai SQLite in-memory (`gorm.io/driver/sqlite` sudah tersedia di go.mod) dengan `db.AutoMigrate` pada entity yang relevan.
- Saat ini ada perubahan lokal user yang belum di-commit di `client/src/routes/program-kerja/[slug]/+page.server.ts` — jangan di-revert, task client hanya mengganti URL API di dalamnya.

---

### Task 1: Repository — `ProgramBySlugOnly` + interface

**Files:**
- Modify: `server/modules/cabinet/repository/repository.go:32` (interface) dan sisipkan implementasi setelah method `ProgramBySlug` (baris ~296)
- Create: `server/modules/cabinet/repository/repository_test.go`

**Interfaces:**
- Produces: `func (r *cabinetRepository) ProgramBySlugOnly(programSlug string, public bool) (*entities.WorkProgram, error)` dan deklarasi pada interface `Repository`. Semua task service/controller memakai ini.

- [ ] **Step 1: Tulis failing test**

`server/modules/cabinet/repository/repository_test.go`:

```go
package repository

import (
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
	return db
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
```

- [ ] **Step 2: Jalankan test, pastikan gagal**

Run: `cd server && go test ./modules/cabinet/repository/ -run TestProgramBySlugOnly -v`
Expected: FAIL compile — `r.ProgramBySlugOnly undefined`.

- [ ] **Step 3: Implementasi repository**

Tambahkan di interface `Repository` (setelah `ProgramBySlug`):

```go
	ProgramBySlugOnly(programSlug string, public bool) (*entities.WorkProgram, error)
```

Tambahkan implementasi (setelah method `ProgramBySlug`):

```go
func (r *cabinetRepository) ProgramBySlugOnly(programSlug string, public bool) (*entities.WorkProgram, error) {
	query := r.db.Joins("JOIN ministries ON ministries.id = work_programs.ministry_id").Joins("JOIN cabinet_terms ON cabinet_terms.id = ministries.cabinet_term_id").Where("work_programs.slug = ?", programSlug)
	if public {
		query = query.Where("ministries.is_active = ? AND ministries.is_published = ? AND cabinet_terms.is_active = ? AND cabinet_terms.is_published = ? AND work_programs.is_published = ?", true, true, true, true, true)
	}
	var value entities.WorkProgram
	err := query.Preload("Ministry").Preload("CoverMedia").Preload("Milestones", func(db *gorm.DB) *gorm.DB { return db.Order("display_order ASC, id ASC") }).Preload("Documentations", func(db *gorm.DB) *gorm.DB { return db.Order("display_order ASC, id ASC") }).Preload("Documentations.MediaAsset").First(&value).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &value, err
}
```

- [ ] **Step 4: Jalankan test, pastikan pass**

Run: `cd server && go test ./modules/cabinet/repository/ -v`
Expected: semua `TestProgramBySlugOnly*` PASS (dan tidak ada test lain di package ini).

- [ ] **Step 5: gofmt + validasi + commit**

```bash
gofmt -l server/modules/cabinet/repository/
cd server && go test ./... -v 2>&1 | tail -30
git add server/modules/cabinet/repository/repository.go server/modules/cabinet/repository/repository_test.go
git commit -m "feat(cabinet): add ProgramBySlugOnly repository method with tests"
```

---

### Task 2: Service — `PublicProgramBySlug` + slug datetime + validasi unik

**Files:**
- Modify: `server/modules/cabinet/service/service.go` (interface baris ~25, implementasi baru setelah `PublicProgram` baris ~116, `CreateProgram` baris ~364, `UpdateProgram` baris ~379, helper baru di dekat `validateProgram`)
- Create: `server/modules/cabinet/service/service_test.go`

**Interfaces:**
- Consumes: `repository.Repository.ProgramBySlugOnly(programSlug string, public bool) (*entities.WorkProgram, error)` (Task 1).
- Produces: `func (s *cabinetService) PublicProgramBySlug(programSlug string) (*dto.ProgramResponse, error)` — dipakai controller Task 3. Perubahan perilaku: `CreateProgram` otomatis menambah suffix datetime; `UpdateProgram` mempertahankan slug lama; duplikat slug manual ditolak.

- [ ] **Step 1: Tulis failing test**

`server/modules/cabinet/service/service_test.go`:

```go
package service

import (
	"regexp"
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
	return db
}

func seedMinistry(t *testing.T, db *gorm.DB) *entities.Ministry {
	t.Helper()
	cabinet := entities.CabinetTerm{Name: "Kabinet Cerita Loka", Slug: "cerita-loka", IsActive: true, IsPublished: true}
	if err := db.Create(&cabinet).Error; err != nil {
		t.Fatal(err)
	}
	ministry := entities.Ministry{Code: "KEM1", Name: "Kementerian Satu", Slug: "kementerian-satu", CabinetTermID: &cabinet.ID, UnitType: "KEMENTERIAN", IsActive: true, IsPublished: true}
	if err := db.Create(&ministry).Error; err != nil {
		t.Fatal(err)
	}
	return &ministry
}

var medinfoActor = &middlewares.Claims{Role: constants.RoleAdminMedinfo}

func TestPublicProgramBySlug(t *testing.T) {
	db := newTestDB(t, "svc_public")
	ministry := seedMinistry(t, db)
	if err := db.Create(&entities.WorkProgram{MinistryID: ministry.ID, Name: "Cerita Visual", Slug: "cerita-visual", LifecycleStatus: "ONGOING", IsPublished: true}).Error; err != nil {
		t.Fatal(err)
	}
	svc := New(repository.New(db))

	value, err := svc.PublicProgramBySlug("cerita-visual")
	if err != nil || value == nil {
		t.Fatalf("expected program, got %v err=%v", value, err)
	}
	if value.Name != "Cerita Visual" || value.MinistryName != "Kementerian Satu" {
		t.Fatalf("unexpected response: %+v", value)
	}
}

func TestPublicProgramBySlugUnpublished(t *testing.T) {
	db := newTestDB(t, "svc_public_unpub")
	ministry := seedMinistry(t, db)
	if err := db.Create(&entities.WorkProgram{MinistryID: ministry.ID, Name: "Cerita Visual", Slug: "cerita-visual", LifecycleStatus: "DRAFT", IsPublished: false}).Error; err != nil {
		t.Fatal(err)
	}
	svc := New(repository.New(db))

	value, err := svc.PublicProgramBySlug("cerita-visual")
	if err != nil || value != nil {
		t.Fatalf("expected nil,nil got %v err=%v", value, err)
	}
}

func TestCreateProgramAutoSlugAppendsDatetime(t *testing.T) {
	db := newTestDB(t, "svc_create_slug")
	ministry := seedMinistry(t, db)
	svc := New(repository.New(db))

	value, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: ministry.ID, Name: "Cerita Visual", LifecycleStatus: "ONGOING"}, medinfoActor)
	if err != nil {
		t.Fatal(err)
	}
	matched, _ := regexp.MatchString(`^cerita-visual-\d{8}-\d{6}$`, value.Slug)
	if !matched {
		t.Fatalf("slug %q does not match cerita-visual-YYYYMMDD-HHMMSS", value.Slug)
	}
}

func TestCreateProgramDuplicateManualSlug(t *testing.T) {
	db := newTestDB(t, "svc_dup_slug")
	ministry := seedMinistry(t, db)
	svc := New(repository.New(db))

	if _, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: ministry.ID, Name: "Satu", Slug: "satu", LifecycleStatus: "ONGOING"}, medinfoActor); err != nil {
		t.Fatal(err)
	}
	_, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: ministry.ID, Name: "Dua", Slug: "satu", LifecycleStatus: "ONGOING"}, medinfoActor)
	if err == nil || err.Error() != "slug program sudah digunakan" {
		t.Fatalf("expected duplicate slug error, got %v", err)
	}
}

func TestUpdateProgramPreservesSlug(t *testing.T) {
	db := newTestDB(t, "svc_update_slug")
	ministry := seedMinistry(t, db)
	svc := New(repository.New(db))

	created, err := svc.CreateProgram(dto.ProgramRequest{MinistryID: ministry.ID, Name: "Cerita Visual", Slug: "cerita-visual", LifecycleStatus: "ONGOING"}, medinfoActor)
	if err != nil {
		t.Fatal(err)
	}
	updated, err := svc.UpdateProgram(created.ID, dto.ProgramRequest{MinistryID: ministry.ID, Name: "Cerita Visual Baru", LifecycleStatus: "COMPLETED"}, medinfoActor)
	if err != nil {
		t.Fatal(err)
	}
	if updated.Slug != "cerita-visual" {
		t.Fatalf("slug changed to %q, expected cerita-visual", updated.Slug)
	}
}
```

- [ ] **Step 2: Jalankan test, pastikan gagal**

Run: `cd server && go test ./modules/cabinet/service/ -v`
Expected: FAIL compile — `svc.PublicProgramBySlug undefined`.

- [ ] **Step 3: Implementasi service**

Tambah di interface `Service` (setelah `PublicProgram`):

```go
	PublicProgramBySlug(programSlug string) (*dto.ProgramResponse, error)
```

Tambah implementasi (setelah method `PublicProgram`):

```go
func (s *cabinetService) PublicProgramBySlug(programSlug string) (*dto.ProgramResponse, error) {
	program, err := s.repo.ProgramBySlugOnly(programSlug, true)
	if err != nil || program == nil {
		return nil, err
	}
	result := programToDTO(*program)
	return &result, nil
}
```

Modifikasi `CreateProgram` — sisipkan dua blok setelah `if err := validateProgram(req); err != nil { return nil, err }` dan sebelum `programFromRequest`:

```go
	if strings.TrimSpace(req.Slug) == "" {
		req.Slug = utils.Slugify(req.Name) + "-" + time.Now().Format("20060102-150405")
	}
	if err := s.validateProgramSlugUnique(req.Slug, 0); err != nil {
		return nil, err
	}
```

Modifikasi `UpdateProgram` — sisipkan setelah `if !s.canManage(actor, value.MinistryID) { return nil, errors.New("akses unit ditolak") }` dan sebelum `validateProgram`:

```go
	if strings.TrimSpace(req.Slug) == "" {
		req.Slug = value.Slug
	}
	if err := s.validateProgramSlugUnique(req.Slug, id); err != nil {
		return nil, err
	}
```

Tambah helper baru (di dekat `validateProgram`):

```go
func (s *cabinetService) validateProgramSlugUnique(slug string, excludeID uint64) error {
	var count int64
	if err := s.repo.DB().Model(&entities.WorkProgram{}).Where("slug = ? AND id <> ?", slug, excludeID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("slug program sudah digunakan")
	}
	return nil
}
```

- [ ] **Step 4: Jalankan test, pastikan pass**

Run: `cd server && go test ./modules/cabinet/service/ -v`
Expected: 5 test PASS (`TestPublicProgramBySlug`, `TestPublicProgramBySlugUnpublished`, `TestCreateProgramAutoSlugAppendsDatetime`, `TestCreateProgramDuplicateManualSlug`, `TestUpdateProgramPreservesSlug`).

- [ ] **Step 5: gofmt + validasi + commit**

```bash
gofmt -l server/modules/cabinet/service/
cd server && go test ./... -v 2>&1 | tail -30
git add server/modules/cabinet/service/service.go server/modules/cabinet/service/service_test.go
git commit -m "feat(cabinet): find program by slug only, datetime-slug on create, preserve slug on update"
```

---

### Task 3: Controller handler + route publik

**Files:**
- Modify: `server/modules/cabinet/controller/controller.go` (tambah handler setelah `PublicProgram` baris ~58)
- Modify: `server/modules/cabinet/routes.go:19` (tambah route di group `public`)
- Create: `server/modules/cabinet/controller/controller_test.go`

**Interfaces:**
- Consumes: `service.Service.PublicProgramBySlug(programSlug string) (*dto.ProgramResponse, error)` (Task 2).
- Produces: route `GET /api/v1/cabinet/programs/:programSlug` (handler `PublicProgramBySlug`) — dipakai client Task 4.

- [ ] **Step 1: Tulis failing test**

`server/modules/cabinet/controller/controller_test.go`:

```go
package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/cabinet/repository"
	"bemunair2026/server/modules/cabinet/service"
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
	return db
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
	if err := db.Create(&entities.WorkProgram{MinistryID: ministry.ID, Name: "Cerita Visual", Slug: "cerita-visual", LifecycleStatus: "ONGOING", IsPublished: true}).Error; err != nil {
		t.Fatal(err)
	}
}

func TestPublicProgramBySlugEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := newTestDB(t, "ctrl_program")
	seedPublishedProgram(t, db)

	router := gin.New()
	router.GET("/cabinet/programs/:programSlug", New(service.New(repository.New(db))).PublicProgramBySlug)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/cabinet/programs/cerita-visual", nil)
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
	}
	var body struct {
		Status bool `json:"status"`
		Data   struct {
			Name string `json:"name"`
		} `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if !body.Status || body.Data.Name != "Cerita Visual" {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/cabinet/programs/tidak-ada", nil)
	router.ServeHTTP(w2, req2)
	if w2.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w2.Code)
	}
}
```

- [ ] **Step 2: Jalankan test, pastikan gagal**

Run: `cd server && go test ./modules/cabinet/controller/ -v`
Expected: FAIL compile — handler tidak terdaftar / undefined.

- [ ] **Step 3: Implementasi controller + route**

Di `controller.go`, tambahkan handler setelah `PublicProgram`:

```go
func (c *Controller) PublicProgramBySlug(ctx *gin.Context) {
	value, err := c.service.PublicProgramBySlug(ctx.Param("programSlug"))
	if err != nil || value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Program kerja tidak ditemukan")
		return
	}
	response.OK(ctx, "Detail program kerja", value)
}
```

Di `routes.go`, pada group `public` (setelah baris `public.GET("/units/:slug/programs/:programSlug", c.PublicProgram)`):

```go
	public.GET("/programs/:programSlug", c.PublicProgramBySlug)
```

- [ ] **Step 4: Jalankan test, pastikan pass**

Run: `cd server && go test ./modules/cabinet/controller/ -v`
Expected: `TestPublicProgramBySlugEndpoint` PASS (200 + 404 case).

- [ ] **Step 5: gofmt + validasi + commit**

```bash
gofmt -l server/modules/cabinet/controller/controller.go
cd server && go test ./... -v 2>&1 | tail -30
git add server/modules/cabinet/controller/controller.go server/modules/cabinet/controller/controller_test.go server/modules/cabinet/routes.go
git commit -m "feat(cabinet): add public endpoint GET /cabinet/programs/:programSlug"
```

---

### Task 4: Client — halaman `program-kerja/[slug]`

**Files:**
- Modify: `client/src/routes/program-kerja/[slug]/+page.server.ts:7`

**Interfaces:**
- Consumes: route `GET /api/v1/cabinet/programs/:programSlug` (Task 3), respon DTO `WorkProgram` yang sama dengan endpoint lama.

- [ ] **Step 1: Ubah URL API**

Baca file dulu (ada perubahan lokal user yang belum di-commit — jangan di-revert). Ganti baris pemanggilan API:

```ts
  export const load: PageServerLoad = async ({ fetch, params }) => {
    try { return { program: (await apiRequest<WorkProgram>(fetch, undefined, `/cabinet/programs/${params.slug}`)).data }; } catch { error(404, 'Program kerja tidak ditemukan'); }
  };
```

- [ ] **Step 2: Validasi client**

Run: `cd client && pnpm check`
Expected: PASS (tanpa error tipe baru).

- [ ] **Step 3: Commit**

```bash
git add client/src/routes/program-kerja/[slug]/+page.server.ts
git commit -m "fix(client): call /cabinet/programs/:slug for program detail page"
```

---

### Task 5: Dokumentasi + validasi akhir

**Files:**
- Modify: `docs/api/cabinet.md`

- [ ] **Step 1: Tambah dokumentasi endpoint baru**

Di `docs/api/cabinet.md`, sisipkan section baru di antara `### GET /api/v1/cabinet/units/:slug/programs/:programSlug` (baris ~31-35) dan `## Response Shape`:

```markdown
### GET /api/v1/cabinet/programs/:programSlug

Mengambil detail program kerja yang dipublikasikan berdasarkan slug program saja (tanpa slug unit), termasuk milestone dan dokumentasi yang aktif. Endpoint lama `GET /api/v1/cabinet/units/:slug/programs/:programSlug` tetap tersedia.

Program tanpa dokumentasi tetap mengembalikan `documentations: []`; frontend menampilkan empty state "Coming Soon".
```

Perbarui baris "Slug program kerja memiliki slug unik dalam satu unit." menjadi:

```markdown
- Slug program otomatis dibentuk dari nama program plus timestamp `YYYYMMDD-HHMMSS` (contoh: `cerita-visual-20260807-153045`) sehingga unik global; slug manual divalidasi tidak boleh duplikat.
```

- [ ] **Step 2: Validasi menyeluruh**

```bash
gofmt -l server/modules/cabinet/
cd server && go test ./... -v 2>&1 | tail -40
cd client && pnpm check && pnpm lint && pnpm build
```

Expected: semua test PASS, check/lint/build PASS. Jika build client gagal karena dependency/daemon eksternal, laporkan perintah yang gagal dan penyebabnya.

- [ ] **Step 3: Uji manual (opsional, bila stack jalan)**

```bash
# dengan stack dev jalan (docker compose -f infra/docker-compose.dev.yml up --build)
curl -s http://localhost:8081/api/v1/cabinet/programs/cerita-visual | head -c 500   # 200
curl -s -o /dev/null -w "%{http_code}\n" http://localhost:8081/api/v1/cabinet/programs/tidak-ada   # 404
```

- [ ] **Step 4: Commit**

```bash
git add docs/api/cabinet.md
git commit -m "docs(api): document GET /cabinet/programs/:programSlug and new slug rule"
```

---

## Self-Review

- **Spec coverage:** repository `ProgramBySlugOnly` (Task 1) ✓; service `PublicProgramBySlug` + slug datetime + preserve slug + validasi unik (Task 2) ✓; controller + route baru, endpoint lama tetap (Task 3) ✓; client URL baru (Task 4) ✓; docs + catatan slug (Task 5) ✓; validasi gofmt/go test/pnpm check per AGENTS.md (tiap task) ✓.
- **Placeholder scan:** semua step berisi kode nyata; tidak ada TBD/TODO.
- **Type consistency:** `ProgramBySlugOnly(programSlug string, public bool) (*entities.WorkProgram, error)` sama di Task 1/2; `PublicProgramBySlug(programSlug string) (*dto.ProgramResponse, error)` sama di Task 2/3; `validateProgramSlugUnique(slug string, excludeID uint64) error` konsisten dipanggil `(req.Slug, 0)` dan `(req.Slug, id)`; pesan error `"slug program sudah digunakan"` sama di implementasi dan test.

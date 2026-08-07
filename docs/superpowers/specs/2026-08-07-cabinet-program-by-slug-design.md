# Design: Find Program Kerja by Program Slug Saja

Tanggal: 2026-08-07
Status: Approved by user (desain disetujui 2026-08-07)
Scope: `server/modules/cabinet`, halaman publik `client/src/routes/program-kerja/[slug]`, `docs/api/cabinet.md`

## 1. Konteks dan Masalah

Halaman publik detail program kerja (`/program-kerja/[slug]`) memanggil:

```
GET /api/v1/cabinet/units/:slug/programs/:programSlug
```

Endpoint ini butuh DUA slug (unit + program), tetapi halaman hanya punya SATU slug (program slug — lihat `ProkerCard.svelte:14` dan `calendar-utils.ts:234` yang membangun URL `/program-kerja/{programSlug}`). Load halaman (`program-kerja/[slug]/+page.server.ts:7`) memakai `params.slug` untuk kedua parameter, sehingga query join `ministries.slug = work_programs.slug` hampir selalu gagal dan halaman selalu 404.

Keputusan user: tambahkan endpoint baru yang cukup find by **program slug saja**; endpoint lama disisakan (tidak dihapus).

## 2. Keputusan Desain (disetujui user)

1. Endpoint baru `GET /api/v1/cabinet/programs/:programSlug` (publik, tanpa auth). Endpoint lama `GET /api/v1/cabinet/units/:slug/programs/:programSlug` TIDAK dihapus.
2. Slug otomatis saat create program: `Slugify(name) + "-" + time.Now().Format("20060102-150405")` (format `YYYYMMDD-HHMMSS`, contoh: `cerita-visual-20260807-153045`). Slug manual dari request dipakai apa adanya.
3. Saat update program dengan `req.Slug` kosong: pertahankan slug lama (tidak regenerasi) — memperbaiki perilaku lama yang bisa mengubah slug dan merusak URL publik.
4. Validasi keunikan slug di service (`validateProgramSlugUnique`) sebagai defense-in-depth untuk slug manual; otomatis unik untuk slug ber-datetime.
5. `WorkProgramCard.svelte` (komponen tidak terpakai, link rusak) KELUAR dari scope.

## 3. Perubahan Server

### 3.1 Repository — `server/modules/cabinet/repository/repository.go`

Tambah method baru (method lama `ProgramBySlug` tidak diubah):

```go
func (r *cabinetRepository) ProgramBySlugOnly(programSlug string, public bool) (*entities.WorkProgram, error)
```

- `Joins("JOIN ministries ON ministries.id = work_programs.ministry_id")`
- `Joins("JOIN cabinet_terms ON cabinet_terms.id = ministries.cabinet_term_id")`
- `Where("work_programs.slug = ?", programSlug)`
- Jika `public`: filter `ministries.is_active AND ministries.is_published AND cabinet_terms.is_active AND cabinet_terms.is_published AND work_programs.is_published` (semua `true`), sama seperti `ProgramBySlug` lama.
- Preload identik dengan `ProgramBySlug`: `Ministry`, `CoverMedia`, `Milestones` (order `display_order ASC, id ASC`), `Documentations` (order sama), `Documentations.MediaAsset`.
- `gorm.ErrRecordNotFound` → return `nil, nil`.

Tambahkan deklarasi method pada interface `Repository`.

### 3.2 Service — `server/modules/cabinet/service/service.go`

- Tambah method interface + implementasi:

```go
func (s *cabinetService) PublicProgramBySlug(programSlug string) (*dto.ProgramResponse, error)
```

  Implementasi: `repo.ProgramBySlugOnly(programSlug, true)` → `programToDTO`; nil/err diteruskan.

- Pembuatan slug saat CREATE (`CreateProgram`), sebelum memanggil `programFromRequest`:

```go
if strings.TrimSpace(req.Slug) == "" {
    req.Slug = utils.Slugify(req.Name) + "-" + time.Now().Format("20060102-150405")
}
```

- Saat UPDATE (`UpdateProgram`), sebelum `programFromRequest`:

```go
if strings.TrimSpace(req.Slug) == "" {
    req.Slug = value.Slug // pertahankan slug lama
}
```

  (Catatan: `programFromRequest` lama meregenerasi slug dari nama saat kosong; dengan ini perilaku regenerasi saat update dihilangkan.)

- Validasi keunikan:

```go
func (s *cabinetService) validateProgramSlugUnique(slug string, excludeID uint64) error
```

  Count `work_programs` di mana `slug = ? AND id <> ?`; jika > 0 → `errors.New("slug program sudah digunakan")`. Dipanggil di `CreateProgram` (excludeID = 0) dan `UpdateProgram` (excludeID = id), setelah slug final ditentukan.

### 3.3 Controller — `server/modules/cabinet/controller/controller.go`

Tambah handler (pola sama dengan `PublicProgram`):

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

### 3.4 Routes — `server/modules/cabinet/routes.go`

Pada group publik `api.Group("/cabinet")`:

```go
public.GET("/programs/:programSlug", c.PublicProgramBySlug)
```

Tidak bentrok dengan `/:slug` (CabinetBySlug) karena `/programs/:programSlug` dua segmen.

## 4. Perubahan Client

`client/src/routes/program-kerja/[slug]/+page.server.ts` — ganti URL API:

```ts
`/cabinet/programs/${params.slug}`
```

`ProkerCard.svelte`, `calendar-utils.ts`, `ListProker.svelte` TIDAK berubah.

## 5. Perubahan Dokumentasi

`docs/api/cabinet.md`:

- Tambah section `### GET /api/v1/cabinet/programs/:programSlug` (ambil detail program publik berdasarkan slug program saja; termasuk milestone dan dokumentasi aktif; program tanpa dokumentasi tetap `documentations: []`).
- Catatan slug diperbarui: slug otomatis kini unik global berkat suffix datetime; slug manual divalidasi unik di service.
- Catatan bahwa endpoint dua-slug lama tetap tersedia.

## 6. Error Handling

- Program tidak ditemukan / err query → `404` envelope `response.Error` "Program kerja tidak ditemukan" (sama dengan endpoint lama).
- Duplikat slug manual → `422` `response.ValidationError` dengan pesan error service.

## 7. Validasi (sesuai AGENTS.md)

- `gofmt` file Go yang berubah.
- `cd server && go test ./... -v` (tidak ada test cabinet saat ini; minimal build + test suite existing).
- `cd client && pnpm check` (+ `pnpm lint && pnpm build` bila memungkinkan).
- Uji manual curl: `GET /api/v1/cabinet/programs/{slug}` → 200 untuk slug yang dipublikasikan, 404 untuk slug tidak ada/unpublished.

## 8. Risiko dan Catatan

- Data lama dengan slug duplikat antar unit: endpoint baru mengembalikan baris pertama (`First()`); risiko kecil, dicatat, tidak dimigrasi di scope ini.
- Perilaku update slug diubah: slug tidak lagi regenerasi dari nama saat update tanpa slug eksplisit — ini intent desain (stabilitas URL).
- Tanpa test otomatis untuk modul cabinet; verifikasi via build + curl manual.

## 9. Di Luar Scope

- `WorkProgramCard.svelte` (komponen tidak terpakai).
- Migration/unique index DB untuk slug.
- Penghapusan endpoint lama.

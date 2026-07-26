package validation

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"bemunair2026/server/modules/content_submission/dto"
	"bemunair2026/server/pkg/constants"
)

type ContentSubmissionValidation struct{}

func NewContentSubmissionValidation() *ContentSubmissionValidation {
	return &ContentSubmissionValidation{}
}

func (v *ContentSubmissionValidation) ValidateCreateRequest(req dto.CreateRequest) error {
	if !isValidType(req.SubmissionType) {
		return errors.New("submission_type harus FEED_INSTAGRAM, REELS_INSTAGRAM, INSTASTORY, atau ARTIKEL")
	}
	title := strings.TrimSpace(req.Title)
	if len([]rune(title)) < 5 || len([]rune(title)) > 150 {
		return errors.New("title harus terdiri dari 5 sampai 150 karakter")
	}
	if strings.TrimSpace(req.Caption) == "" {
		return errors.New("caption wajib diisi")
	}
	if req.PublishDate == nil {
		return errors.New("publish_date wajib diisi")
	}
	if isBlank(req.PublishTime) {
		return errors.New("publish_time wajib diisi")
	}
	if err := validatePublishTime(*req.PublishTime); err != nil {
		return err
	}

	switch req.SubmissionType {
	case constants.ContentTypeFeedsReels, constants.ContentTypeFeed, constants.ContentTypeInstastory:
		if req.BriefLink == "" {
			return errors.New("brief konten wajib diunggah")
		}
		if isBlank(req.DesignDriveLink) {
			return errors.New("gambar final wajib diunggah")
		}
		if isBlank(req.CanvaLink) || !hasHost(*req.CanvaLink, "canva.com") {
			return errors.New("link Canva yang valid wajib diisi untuk Feed dan Instastory")
		}
		if req.PublishDate.Before(startOfDay(time.Now()).AddDate(0, 0, 7)) {
			return errors.New("tanggal publikasi konten minimal tujuh hari dari hari ini")
		}
	case constants.ContentTypeReels:
		if req.BriefLink == "" {
			return errors.New("brief konten wajib diunggah")
		}
		if isBlank(req.DesignDriveLink) || !isGoogleDrive(*req.DesignDriveLink) {
			return errors.New("link Google Drive video final wajib diisi untuk Reels")
		}
		if req.PublishDate.Before(startOfDay(time.Now()).AddDate(0, 0, 7)) {
			return errors.New("tanggal publikasi konten minimal tujuh hari dari hari ini")
		}
	case constants.ContentTypeArtikel:
		if req.PublishDate.Before(startOfDay(time.Now()).AddDate(0, 0, 3)) {
			return errors.New("tanggal publikasi artikel minimal tiga hari dari hari ini")
		}
		if isBlank(req.DocumentationDriveLink) || !isGoogleDrive(*req.DocumentationDriveLink) {
			return errors.New("link Google Drive dokumentasi wajib diisi untuk artikel")
		}
		if isBlank(req.ArticleDriveLink) || !hasHost(*req.ArticleDriveLink, "docs.google.com") {
			return errors.New("link Google Docs isi artikel wajib diisi")
		}
	}
	return nil
}

func validatePublishTime(value string) error {
	parsed, err := time.Parse("15:04", value)
	if err != nil {
		return errors.New("publish_time harus menggunakan format HH:mm")
	}
	minutes := parsed.Hour()*60 + parsed.Minute()
	if minutes < 8*60 || minutes > 17*60 {
		return errors.New("publish_time harus berada antara 08:00 dan 17:00")
	}
	if parsed.Minute()%30 != 0 {
		return errors.New("publish_time harus menggunakan interval 30 menit")
	}
	return nil
}

func isGoogleDrive(raw string) bool {
	return hasHost(raw, "drive.google.com") || hasHost(raw, "docs.google.com")
}

func hasHost(raw, expected string) bool {
	parsed, err := url.ParseRequestURI(strings.TrimSpace(raw))
	if err != nil {
		return false
	}
	host := strings.ToLower(parsed.Hostname())
	return host == expected || strings.HasSuffix(host, "."+expected)
}

func isValidType(value string) bool {
	return value == constants.ContentTypeFeedsReels ||
		value == constants.ContentTypeFeed ||
		value == constants.ContentTypeReels ||
		value == constants.ContentTypeInstastory ||
		value == constants.ContentTypeArtikel
}

func startOfDay(value time.Time) time.Time {
	return time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location())
}

func isBlank(value *string) bool {
	return value == nil || strings.TrimSpace(*value) == ""
}

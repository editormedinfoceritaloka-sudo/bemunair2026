package validation

import (
	"errors"

	"bemunair2026/server/modules/content_submission/dto"
	"bemunair2026/server/pkg/constants"
)

type ContentSubmissionValidation struct{}

func NewContentSubmissionValidation() *ContentSubmissionValidation {
	return &ContentSubmissionValidation{}
}

func (v *ContentSubmissionValidation) ValidateCreateRequest(req dto.CreateRequest) error {
	if !isValidType(req.SubmissionType) {
		return errors.New("submission_type harus salah satu dari FEEDS_REELS, INSTASTORY, ARTIKEL")
	}
	if req.Title == "" {
		return errors.New("title wajib diisi")
	}
	if req.Caption == "" {
		return errors.New("caption wajib diisi")
	}
	if req.BriefLink == "" {
		return errors.New("brief_link wajib diisi")
	}

	switch req.SubmissionType {
	case constants.ContentTypeFeedsReels, constants.ContentTypeInstastory:
		if req.PublishDate == nil {
			return errors.New("publish_date wajib diisi untuk Feeds & Reels / Instastory")
		}
		if isBlank(req.PublishTime) {
			return errors.New("publish_time wajib diisi untuk Feeds & Reels / Instastory")
		}
		if isBlank(req.DesignDriveLink) {
			return errors.New("design_drive_link wajib diisi untuk Feeds & Reels / Instastory")
		}
		if isBlank(req.CanvaLink) {
			return errors.New("canva_link wajib diisi untuk Feeds & Reels / Instastory")
		}
	case constants.ContentTypeArtikel:
		if isBlank(req.ArticleDriveLink) {
			return errors.New("article_drive_link wajib diisi untuk Artikel")
		}
	}

	return nil
}

func isValidType(t string) bool {
	return t == constants.ContentTypeFeedsReels ||
		t == constants.ContentTypeInstastory ||
		t == constants.ContentTypeArtikel
}

func isBlank(s *string) bool {
	return s == nil || *s == ""
}

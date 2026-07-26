package validation

import (
	"strings"
	"testing"
	"time"

	"bemunair2026/server/modules/content_submission/dto"
	"bemunair2026/server/pkg/constants"
)

func TestValidateReelsUsesDriveWithoutMediaUpload(t *testing.T) {
	date := startOfDay(time.Now()).AddDate(0, 0, 8)
	phoneTime := "10:30"
	drive := "https://drive.google.com/file/d/video/view"
	briefID := "brief-file"
	briefMime := "application/pdf"
	req := dto.CreateRequest{
		SubmissionType:    constants.ContentTypeReels,
		Title:             "Reels kegiatan kementerian",
		Caption:           "Caption pengajuan reels.",
		PublishDate:       &date,
		PublishTime:       &phoneTime,
		DesignDriveLink:   &drive,
		BriefFileID:       &briefID,
		BriefFileMimeType: &briefMime,
		BriefFileSize:     1024,
	}
	if err := NewContentSubmissionValidation().ValidateCreateRequest(req); err != nil {
		t.Fatalf("valid Reels rejected: %v", err)
	}
}

func TestValidateFeedRequiresImageMetadataAndCanva(t *testing.T) {
	date := startOfDay(time.Now()).AddDate(0, 0, 8)
	publishTime := "09:00"
	canva := "https://www.canva.com/design/example"
	imageID, imageMime := "image-file", "image/png"
	briefID, briefMime := "brief-file", "application/pdf"
	req := dto.CreateRequest{
		SubmissionType:    constants.ContentTypeFeed,
		Title:             "Feed informasi kegiatan",
		Caption:           "Caption pengajuan feed.",
		PublishDate:       &date,
		PublishTime:       &publishTime,
		CanvaLink:         &canva,
		MediaFileID:       &imageID,
		MediaFileMimeType: &imageMime,
		MediaFileSize:     1024,
		BriefFileID:       &briefID,
		BriefFileMimeType: &briefMime,
		BriefFileSize:     1024,
	}
	if err := NewContentSubmissionValidation().ValidateCreateRequest(req); err != nil {
		t.Fatalf("valid Feed rejected: %v", err)
	}
	req.MediaFileSize = 21 * 1024 * 1024
	if err := NewContentSubmissionValidation().ValidateCreateRequest(req); err == nil || !strings.Contains(err.Error(), "20 MB") {
		t.Fatalf("expected image size validation, got %v", err)
	}
}

func TestValidateArticleMinimumThreeDaysAndGoogleLinks(t *testing.T) {
	date := startOfDay(time.Now()).AddDate(0, 0, 3)
	publishTime := "13:30"
	documentation := "https://drive.google.com/drive/folders/example"
	article := "https://docs.google.com/document/d/example/edit"
	req := dto.CreateRequest{
		SubmissionType:         constants.ContentTypeArtikel,
		Title:                  "Liputan kegiatan mahasiswa",
		Caption:                "Caption artikel kegiatan.",
		PublishDate:            &date,
		PublishTime:            &publishTime,
		DocumentationDriveLink: &documentation,
		ArticleDriveLink:       &article,
		BriefLink:              "-",
	}
	if err := NewContentSubmissionValidation().ValidateCreateRequest(req); err != nil {
		t.Fatalf("valid Article rejected: %v", err)
	}
}

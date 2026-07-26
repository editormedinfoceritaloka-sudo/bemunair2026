package submission_support

import (
	"fmt"
	"html"
	"net/http"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/middlewares"
	"bemunair2026/server/pkg/constants"
	"bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (c *controller) createArticleDraft(ctx *gin.Context) {
	var submission entities.ContentSubmission
	if err := c.db.First(&submission, ctx.Param("id")).Error; err != nil {
		utils.Error(ctx, http.StatusNotFound, utils.NotFound, "Pengajuan artikel tidak ditemukan")
		return
	}
	if submission.ServiceType != constants.ServiceTypeArticle {
		utils.Error(ctx, http.StatusUnprocessableEntity, utils.ValidationError, "Hanya pengajuan artikel yang dapat dibuat menjadi draft")
		return
	}
	if submission.Status != constants.StatusApproved &&
		submission.Status != constants.StatusScheduled &&
		submission.Status != constants.StatusPublished {
		utils.Error(ctx, http.StatusConflict, utils.ValidationError, "Pengajuan harus disetujui sebelum dibuat menjadi draft artikel")
		return
	}

	var existing entities.Article
	result := c.db.Where("source_submission_id = ?", submission.ID).First(&existing)
	if result.Error == nil {
		utils.OK(ctx, "Draft artikel sudah tersedia", gin.H{"id": existing.ID, "created": false})
		return
	}
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		utils.Error(ctx, http.StatusInternalServerError, utils.InternalError, result.Error.Error())
		return
	}

	slug, err := c.uniqueArticleSlug(submission.Title)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, utils.InternalError, err.Error())
		return
	}
	body := fmt.Sprintf(
		"<h1>%s</h1><p>#CeritaHariIni</p><h2>Kepala berita</h2><p>Mulai tulis inti kegiatan di sini.</p><h2>Tubuh berita</h2><p>Kembangkan rangkaian kegiatan, peserta, tujuan, dan hasil.</p><h2>Ekor berita</h2><p>Tutup artikel dengan kesimpulan atau tindak lanjut.</p><p><strong>Sumber naskah:</strong> <a href=\"%s\">Google Docs pengaju</a></p>",
		html.EscapeString(submission.Title),
		html.EscapeString(stringValue(submission.ArticleDriveLink)),
	)
	article := entities.Article{
		Slug: slug, Title: submission.Title, Body: body,
		AuthorID:           middlewares.CurrentClaims(ctx).UserID,
		SourceSubmissionID: &submission.ID,
		Status:             constants.ArticleStatusDraft,
	}
	if err := c.db.Create(&article).Error; err != nil {
		utils.Error(ctx, http.StatusConflict, utils.ValidationError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, utils.BuildResponseSuccess("Draft artikel berhasil dibuat", gin.H{"id": article.ID, "created": true}))
}

func (c *controller) uniqueArticleSlug(title string) (string, error) {
	base := utils.Slugify(title)
	if base == "" {
		base = "artikel"
	}
	for index := 1; ; index++ {
		slug := base
		if index > 1 {
			slug = fmt.Sprintf("%s-%d", base, index)
		}
		var count int64
		if err := c.db.Model(&entities.Article{}).Where("slug = ?", slug).Count(&count).Error; err != nil {
			return "", err
		}
		if count == 0 {
			return slug, nil
		}
	}
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

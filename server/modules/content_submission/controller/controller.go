package controller

import (
	"net/http"
	"strconv"
	"time"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/content_submission/dto"
	"bemunair2026/server/modules/content_submission/service"
	"bemunair2026/server/modules/content_submission/validation"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ContentSubmissionController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Timeline(ctx *gin.Context)
	SubmitRevision(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	AssignPJ(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type contentSubmissionController struct {
	service    service.ContentSubmissionService
	validation *validation.ContentSubmissionValidation
}

var _ ContentSubmissionController = (*contentSubmissionController)(nil)

func NewContentSubmissionController(service service.ContentSubmissionService) ContentSubmissionController {
	return &contentSubmissionController{
		service:    service,
		validation: validation.NewContentSubmissionValidation(),
	}
}

func (c *contentSubmissionController) Create(ctx *gin.Context) {
	claims := middlewares.CurrentClaims(ctx)

	var publishDate *time.Time
	if raw := ctx.PostForm("publish_date"); raw != "" {
		parsed, err := time.ParseInLocation("2006-01-02", raw, time.Local)
		if err != nil {
			res := response.BuildResponseFailed("publish_date harus format YYYY-MM-DD", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
			return
		}
		publishDate = &parsed
	}

	req := dto.CreateRequest{
		Ministry:               ctx.PostForm("ministry"),
		ServiceType:            ctx.PostForm("service_type"),
		ContentFormat:          optionalForm(ctx, "content_format"),
		SubmitterPhone:         optionalForm(ctx, "submitter_phone"),
		SubmissionType:         ctx.PostForm("submission_type"),
		Title:                  ctx.PostForm("title"),
		Caption:                ctx.PostForm("caption"),
		AddSong:                optionalForm(ctx, "add_song"),
		AdditionalNotes:        optionalForm(ctx, "additional_notes"),
		PublishDate:            publishDate,
		PublishTime:            optionalForm(ctx, "publish_time"),
		DesignDriveLink:        optionalForm(ctx, "design_drive_link"),
		CanvaLink:              optionalForm(ctx, "canva_link"),
		ArticleDriveLink:       optionalForm(ctx, "article_drive_link"),
		DocumentationDriveLink: optionalForm(ctx, "documentation_drive_link"),
		RequiredInformation:    optionalForm(ctx, "required_information"),
		SongTitle:              optionalForm(ctx, "song_title"),
		SongArtist:             optionalForm(ctx, "song_artist"),
		MediaFileID:            optionalForm(ctx, "media_file_id"),
		MediaFileName:          optionalForm(ctx, "media_file_name"),
		MediaFileMimeType:      optionalForm(ctx, "media_file_mime_type"),
		MediaFileSize:          uintForm(ctx, "media_file_size"),
		BriefFileID:            optionalForm(ctx, "brief_file_id"),
		BriefFileName:          optionalForm(ctx, "brief_file_name"),
		BriefFileMimeType:      optionalForm(ctx, "brief_file_mime_type"),
		BriefFileSize:          uintForm(ctx, "brief_file_size"),
		BriefLink:              ctx.PostForm("brief_link"),
	}

	if err := c.validation.ValidateCreateRequest(req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	created, warnings, err := c.service.Create(req, claims.UserID, claims.Ministry)
	if err != nil {
		res := response.BuildResponseFailed("Submission gagal dibuat", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	if len(warnings) > 0 {
		gin.DefaultWriter.Write([]byte("WA warning: " + warnings[0].Error() + "\n"))
	}

	res := response.BuildResponseSuccess("Submission berhasil dibuat", created)
	ctx.JSON(http.StatusCreated, res)
}

func (c *contentSubmissionController) List(ctx *gin.Context) {
	claims := middlewares.CurrentClaims(ctx)
	rows, err := c.service.ListForUser(claims.Role, claims.UserID, claims.Ministry)
	if err != nil {
		res := response.BuildResponseFailed("Gagal mengambil submission", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Daftar content submission", rows)
	res.Meta = response.Meta{Page: 1, PerPage: len(rows), Total: int64(len(rows)), TotalPages: 1}
	ctx.JSON(http.StatusOK, res)
}

func (c *contentSubmissionController) Get(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	row, err := c.service.Get(id, middlewares.CurrentClaims(ctx).Role, middlewares.CurrentClaims(ctx).UserID, middlewares.CurrentClaims(ctx).Ministry)
	if err != nil || row == nil {
		res := response.BuildResponseFailed("Submission tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Detail content submission", row)
	ctx.JSON(http.StatusOK, res)
}

func (c *contentSubmissionController) Timeline(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	claims := middlewares.CurrentClaims(ctx)
	rows, err := c.service.Timeline(id, claims.Role, claims.UserID, claims.Ministry)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.BuildResponseFailed("Submission tidak ditemukan", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.BuildResponseSuccess("Timeline submission", rows))
}

func (c *contentSubmissionController) SubmitRevision(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.BuildResponseFailed("Validasi gagal", err.Error(), nil))
		return
	}
	claims := middlewares.CurrentClaims(ctx)
	row, err := c.service.SubmitRevision(id, claims.Role, claims.UserID, claims.Ministry, req.Notes)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, response.BuildResponseFailed("Revisi gagal dikirim", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.BuildResponseSuccess("Revisi berhasil dikirim", row))
}

func (c *contentSubmissionController) UpdateStatus(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Status == "" {
		res := response.BuildResponseFailed("Validasi gagal", response.ValidationError, nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	row, err := c.service.UpdateStatus(id, req, middlewares.CurrentClaims(ctx).UserID)
	if err != nil {
		res := response.BuildResponseFailed("Status gagal diperbarui", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}

	res := response.BuildResponseSuccess("Status berhasil diperbarui", row)
	ctx.JSON(http.StatusOK, res)
}

func (c *contentSubmissionController) AssignPJ(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.AssignPJRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || req.AssignedPJID == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.BuildResponseFailed("Validasi gagal", response.ValidationError, nil))
		return
	}
	row, err := c.service.AssignPJ(id, req.AssignedPJID, middlewares.CurrentClaims(ctx).UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, response.BuildResponseFailed("PJ gagal ditetapkan", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.BuildResponseSuccess("PJ berhasil ditetapkan", row))
}

func (c *contentSubmissionController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(id); err != nil {
		res := response.BuildResponseFailed("Submission gagal dihapus", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Submission berhasil dihapus", nil)
	ctx.JSON(http.StatusOK, res)
}

func uintForm(ctx *gin.Context, key string) uint64 {
	value, _ := strconv.ParseUint(ctx.PostForm(key), 10, 64)
	return value
}

func optionalForm(ctx *gin.Context, key string) *string {
	if v := ctx.PostForm(key); v != "" {
		return &v
	}
	return nil
}

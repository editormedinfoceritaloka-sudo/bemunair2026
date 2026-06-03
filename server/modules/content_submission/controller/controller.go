package controller

import (
	"net/http"
	"strconv"
	"time"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/content_submission/dto"
	"bemunair2026/server/modules/content_submission/service"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ContentSubmissionController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type contentSubmissionController struct {
	service service.ContentSubmissionService
}

var _ ContentSubmissionController = (*contentSubmissionController)(nil)

func NewContentSubmissionController(service service.ContentSubmissionService) ContentSubmissionController {
	return &contentSubmissionController{service: service}
}

func (c *contentSubmissionController) Create(ctx *gin.Context) {
	claims := middlewares.CurrentClaims(ctx)
	deadline, err := time.Parse(time.RFC3339, ctx.PostForm("deadline"))
	if err != nil {
		res := response.BuildResponseFailed("Deadline harus RFC3339", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	req := dto.CreateRequest{
		Ministry:       ctx.PostForm("ministry"),
		Platform:       ctx.PostForm("platform"),
		SubmissionType: ctx.PostForm("submission_type"),
		Caption:        ctx.PostForm("caption"),
		Deadline:       deadline,
	}
	if file, err := ctx.FormFile("brief_file"); err == nil {
		req.BriefFile = file.Filename
	}
	if file, err := ctx.FormFile("poster_file"); err == nil {
		req.PosterFile = file.Filename
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
	row, err := c.service.Get(id)
	if err != nil || row == nil {
		res := response.BuildResponseFailed("Submission tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Detail content submission", row)
	ctx.JSON(http.StatusOK, res)
}

func (c *contentSubmissionController) UpdateStatus(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Status == "" {
		res := response.BuildResponseFailed("Validasi gagal", response.ValidationError, nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	row, err := c.service.UpdateStatus(id, req)
	if err != nil {
		res := response.BuildResponseFailed("Status gagal diperbarui", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}

	res := response.BuildResponseSuccess("Status berhasil diperbarui", row)
	ctx.JSON(http.StatusOK, res)
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

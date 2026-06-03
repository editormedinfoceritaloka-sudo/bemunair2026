package controller

import (
	"net/http"
	"strconv"
	"time"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/letter_submission/dto"
	"bemunair2026/server/modules/letter_submission/service"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type LetterSubmissionController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type letterSubmissionController struct {
	service service.LetterSubmissionService
}

var _ LetterSubmissionController = (*letterSubmissionController)(nil)

func NewLetterSubmissionController(service service.LetterSubmissionService) LetterSubmissionController {
	return &letterSubmissionController{service: service}
}

func (c *letterSubmissionController) Create(ctx *gin.Context) {
	claims := middlewares.CurrentClaims(ctx)
	var req dto.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	deadline, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		res := response.BuildResponseFailed("Deadline harus RFC3339", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}
	req.DeadlineAt = deadline

	created, warnings, err := c.service.Create(req, claims.UserID, claims.Ministry)
	if err != nil {
		res := response.BuildResponseFailed("Submission gagal dibuat", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	if len(warnings) > 0 {
		gin.DefaultWriter.Write([]byte("WA warning: " + warnings[0].Error() + "\n"))
	}

	res := response.BuildResponseSuccess("Submission surat berhasil dibuat", created)
	ctx.JSON(http.StatusCreated, res)
}

func (c *letterSubmissionController) List(ctx *gin.Context) {
	claims := middlewares.CurrentClaims(ctx)
	rows, err := c.service.ListForUser(claims.Role, claims.UserID, claims.Ministry)
	if err != nil {
		res := response.BuildResponseFailed("Gagal mengambil submission", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Daftar letter submission", rows)
	res.Meta = response.Meta{Page: 1, PerPage: len(rows), Total: int64(len(rows)), TotalPages: 1}
	ctx.JSON(http.StatusOK, res)
}

func (c *letterSubmissionController) Get(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	row, err := c.service.Get(id)
	if err != nil || row == nil {
		res := response.BuildResponseFailed("Submission tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Detail letter submission", row)
	ctx.JSON(http.StatusOK, res)
}

func (c *letterSubmissionController) UpdateStatus(ctx *gin.Context) {
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

func (c *letterSubmissionController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(id); err != nil {
		res := response.BuildResponseFailed("Submission gagal dihapus", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Submission berhasil dihapus", nil)
	ctx.JSON(http.StatusOK, res)
}

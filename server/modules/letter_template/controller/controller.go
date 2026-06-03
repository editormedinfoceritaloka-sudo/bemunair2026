package controller

import (
	"net/http"
	"strconv"

	"bemunair2026/server/modules/letter_template/dto"
	"bemunair2026/server/modules/letter_template/service"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type LetterTemplateController interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type letterTemplateController struct {
	service service.LetterTemplateService
}

var _ LetterTemplateController = (*letterTemplateController)(nil)

func NewLetterTemplateController(service service.LetterTemplateService) LetterTemplateController {
	return &letterTemplateController{service: service}
}

func (c *letterTemplateController) Create(ctx *gin.Context) {
	var req dto.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	template, err := c.service.Create(req)
	if err != nil {
		res := response.BuildResponseFailed("Template gagal dibuat", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Template berhasil dibuat", template)
	ctx.JSON(http.StatusCreated, res)
}

func (c *letterTemplateController) List(ctx *gin.Context) {
	rows, err := c.service.List()
	if err != nil {
		res := response.BuildResponseFailed("Gagal mengambil template", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Daftar template", rows)
	res.Meta = response.Meta{Page: 1, PerPage: len(rows), Total: int64(len(rows)), TotalPages: 1}
	ctx.JSON(http.StatusOK, res)
}

func (c *letterTemplateController) Get(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	row, err := c.service.Get(id)
	if err != nil || row == nil {
		res := response.BuildResponseFailed("Template tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Detail template", row)
	ctx.JSON(http.StatusOK, res)
}

func (c *letterTemplateController) Update(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	row, err := c.service.Update(id, req)
	if err != nil || row == nil {
		res := response.BuildResponseFailed("Template gagal diperbarui", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Template berhasil diperbarui", row)
	ctx.JSON(http.StatusOK, res)
}

func (c *letterTemplateController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(id); err != nil {
		res := response.BuildResponseFailed("Template gagal dihapus", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Template berhasil dihapus", nil)
	ctx.JSON(http.StatusOK, res)
}

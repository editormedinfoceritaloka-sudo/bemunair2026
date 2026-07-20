package controller

import (
	"net/http"
	"strconv"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/article/dto"
	"bemunair2026/server/modules/article/service"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	ListPublished(ctx *gin.Context)
	GetBySlug(ctx *gin.Context)
	ListAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	SetPublished(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type articleController struct {
	service service.ArticleService
}

var _ ArticleController = (*articleController)(nil)

func NewArticleController(service service.ArticleService) ArticleController {
	return &articleController{service: service}
}

func (c *articleController) ListPublished(ctx *gin.Context) {
	page, perPage := paging(ctx)
	result, err := c.service.ListPublished(page, perPage)
	if err != nil {
		res := response.BuildResponseFailed("Gagal mengambil artikel", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	writeList(ctx, "Daftar artikel", result)
}

func (c *articleController) GetBySlug(ctx *gin.Context) {
	article, err := c.service.GetPublishedBySlug(ctx.Param("slug"))
	if err != nil || article == nil {
		res := response.BuildResponseFailed("Artikel tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Detail artikel", article)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) ListAll(ctx *gin.Context) {
	page, perPage := paging(ctx)
	result, err := c.service.ListAll(page, perPage)
	if err != nil {
		res := response.BuildResponseFailed("Gagal mengambil artikel", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	writeList(ctx, "Daftar artikel", result)
}

func (c *articleController) GetByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	article, err := c.service.GetByID(id)
	if err != nil || article == nil {
		res := response.BuildResponseFailed("Artikel tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Detail artikel", article)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) Create(ctx *gin.Context) {
	claims := middlewares.CurrentClaims(ctx)
	var req dto.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	article, err := c.service.Create(req, claims.UserID)
	if err != nil {
		res := response.BuildResponseFailed("Artikel gagal dibuat", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := response.BuildResponseSuccess("Artikel berhasil dibuat", article)
	ctx.JSON(http.StatusCreated, res)
}

func (c *articleController) Update(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	article, err := c.service.Update(id, req)
	if err != nil {
		res := response.BuildResponseFailed("Artikel gagal diperbarui", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}
	if article == nil {
		res := response.BuildResponseFailed("Artikel tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Artikel berhasil diperbarui", article)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) SetPublished(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.PublishRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	article, err := c.service.SetPublished(id, req.Published)
	if err != nil || article == nil {
		res := response.BuildResponseFailed("Artikel tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Status artikel berhasil diperbarui", article)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(id); err != nil {
		res := response.BuildResponseFailed("Artikel gagal dihapus", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Artikel berhasil dihapus", nil)
	ctx.JSON(http.StatusOK, res)
}

func paging(ctx *gin.Context) (int, int) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	perPage, _ := strconv.Atoi(ctx.Query("per_page"))
	return page, perPage
}

func writeList(ctx *gin.Context, message string, result *service.ListResult) {
	res := response.BuildResponseSuccess(message, result.Items)
	res.Meta = response.Meta{
		Page:       result.Page,
		PerPage:    result.PerPage,
		Total:      result.Total,
		TotalPages: result.TotalPages,
	}
	ctx.JSON(http.StatusOK, res)
}

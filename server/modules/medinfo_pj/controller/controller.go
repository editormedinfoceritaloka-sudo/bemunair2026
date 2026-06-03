package controller

import (
	"net/http"
	"strconv"

	"bemunair2026/server/modules/medinfo_pj/dto"
	"bemunair2026/server/modules/medinfo_pj/service"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type MedinfoPJController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Reorder(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type medinfoPJController struct {
	service service.MedinfoPJService
}

var _ MedinfoPJController = (*medinfoPJController)(nil)

func NewMedinfoPJController(service service.MedinfoPJService) MedinfoPJController {
	return &medinfoPJController{service: service}
}

func (c *medinfoPJController) List(ctx *gin.Context) {
	rows, err := c.service.List()
	if err != nil {
		res := response.BuildResponseFailed("Gagal mengambil queue", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Queue PJ Medinfo", rows)
	res.Meta = response.Meta{Page: 1, PerPage: len(rows), Total: int64(len(rows)), TotalPages: 1}
	ctx.JSON(http.StatusOK, res)
}

func (c *medinfoPJController) Create(ctx *gin.Context) {
	var req dto.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || req.UserID == 0 {
		res := response.BuildResponseFailed("Validasi gagal", response.ValidationError, nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	row, err := c.service.Create(req)
	if err != nil {
		res := response.BuildResponseFailed("Queue gagal dibuat", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}

	res := response.BuildResponseSuccess("Queue berhasil dibuat", row)
	ctx.JSON(http.StatusCreated, res)
}

func (c *medinfoPJController) Reorder(ctx *gin.Context) {
	var req dto.ReorderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || len(req.IDs) == 0 {
		res := response.BuildResponseFailed("Validasi gagal", response.ValidationError, nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	if err := c.service.Reorder(req); err != nil {
		res := response.BuildResponseFailed("Queue gagal diurutkan", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Queue berhasil diurutkan", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *medinfoPJController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(id); err != nil {
		res := response.BuildResponseFailed("Queue gagal dihapus", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Queue berhasil dihapus", nil)
	ctx.JSON(http.StatusOK, res)
}

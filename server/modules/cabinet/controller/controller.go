package controller

import (
	"net/http"
	"strconv"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/cabinet/dto"
	"bemunair2026/server/modules/cabinet/service"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Controller struct{ service service.Service }

func New(s service.Service) *Controller { return &Controller{service: s} }

func (c *Controller) ActiveCabinet(ctx *gin.Context) {
	value, err := c.service.PublicCabinet()
	if err != nil || value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Kabinet aktif tidak ditemukan")
		return
	}
	response.OK(ctx, "Kabinet aktif", value)
}
func (c *Controller) CabinetBySlug(ctx *gin.Context) {
	value, err := c.service.PublicCabinetBySlug(ctx.Param("slug"))
	if err != nil || value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Kabinet tidak ditemukan")
		return
	}
	response.OK(ctx, "Detail kabinet", value)
}
func (c *Controller) PublicUnit(ctx *gin.Context) {
	value, err := c.service.PublicUnit(ctx.Param("slug"))
	if err != nil || value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Unit organisasi tidak ditemukan")
		return
	}
	response.OK(ctx, "Detail unit organisasi", value)
}
func (c *Controller) PublicPrograms(ctx *gin.Context) {
	page := queryInt(ctx, "page", 1)
	result, err := c.service.PublicPrograms(ctx.Param("slug"), page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.InternalError, "Gagal mengambil program kerja")
		return
	}
	writePage(ctx, "Program kerja", result.Page, result.PerPage, result.Total, result.TotalPages, result.Items)
}
func (c *Controller) PublicProgram(ctx *gin.Context) {
	value, err := c.service.PublicProgram(ctx.Param("slug"), ctx.Param("programSlug"))
	if err != nil || value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Program kerja tidak ditemukan")
		return
	}
	response.OK(ctx, "Detail program kerja", value)
}
func (c *Controller) PublicProgramBySlug(ctx *gin.Context) {
	value, err := c.service.PublicProgramBySlug(ctx.Param("programSlug"))
	if err != nil || value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Program kerja tidak ditemukan")
		return
	}
	response.OK(ctx, "Detail program kerja", value)
}

func (c *Controller) ListCabinets(ctx *gin.Context) {
	result, err := c.service.Cabinets(queryInt(ctx, "page", 1), queryInt(ctx, "per_page", 20))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.InternalError, "Gagal mengambil kabinet")
		return
	}
	writePage(ctx, "Daftar kabinet", result.Page, result.PerPage, result.Total, result.TotalPages, result.Items)
}
func (c *Controller) CreateCabinet(ctx *gin.Context) {
	var req dto.CabinetRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.CreateCabinet(req)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Kabinet dibuat", value))
}
func (c *Controller) UpdateCabinet(ctx *gin.Context) {
	var req dto.CabinetRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.UpdateCabinet(paramID(ctx), req)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, err.Error())
		return
	}
	if value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Kabinet tidak ditemukan")
		return
	}
	response.OK(ctx, "Kabinet diperbarui", value)
}

func (c *Controller) ListUnits(ctx *gin.Context) {
	cabinetID := uint64(queryInt(ctx, "cabinet_id", 0))
	values, err := c.service.Units(cabinetID, ctx.Query("unit_type"), false)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.InternalError, "Gagal mengambil unit organisasi")
		return
	}
	response.OK(ctx, "Daftar unit organisasi", values)
}
func (c *Controller) CreateUnit(ctx *gin.Context) {
	var req dto.UnitRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.CreateUnit(req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Unit organisasi dibuat", value))
}
func (c *Controller) UpdateUnit(ctx *gin.Context) {
	var req dto.UnitRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.UpdateUnit(paramID(ctx), req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	if value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Unit organisasi tidak ditemukan")
		return
	}
	response.OK(ctx, "Unit organisasi diperbarui", value)
}

func (c *Controller) ListMembers(ctx *gin.Context) {
	values, err := c.service.Members(paramID(ctx), false)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.InternalError, "Gagal mengambil profil pimpinan")
		return
	}
	response.OK(ctx, "Daftar profil pimpinan", values)
}
func (c *Controller) CreateMember(ctx *gin.Context) {
	var req dto.MemberRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.CreateMember(req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Profil pimpinan dibuat", value))
}
func (c *Controller) UpdateMember(ctx *gin.Context) {
	var req dto.MemberRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.UpdateMember(paramID(ctx), req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	if value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Profil pimpinan tidak ditemukan")
		return
	}
	response.OK(ctx, "Profil pimpinan diperbarui", value)
}

func (c *Controller) ListPrograms(ctx *gin.Context) {
	unitID := uint64(queryInt(ctx, "unit_id", 0))
	page := queryInt(ctx, "page", 1)
	perPage := queryInt(ctx, "per_page", 20)
	result, err := c.service.Programs(unitID, page, perPage, false)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, response.InternalError, "Gagal mengambil program kerja")
		return
	}
	writePage(ctx, "Program kerja", result.Page, result.PerPage, result.Total, result.TotalPages, result.Items)
}
func (c *Controller) AdminProgram(ctx *gin.Context) {
	value, err := c.service.Program(paramID(ctx))
	if err != nil || value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Program kerja tidak ditemukan")
		return
	}
	response.OK(ctx, "Detail program kerja", value)
}
func (c *Controller) CreateProgram(ctx *gin.Context) {
	var req dto.ProgramRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.CreateProgram(req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Program kerja dibuat", value))
}
func (c *Controller) UpdateProgram(ctx *gin.Context) {
	var req dto.ProgramRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.UpdateProgram(paramID(ctx), req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	if value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Program kerja tidak ditemukan")
		return
	}
	response.OK(ctx, "Program kerja diperbarui", value)
}
func (c *Controller) PublishProgram(ctx *gin.Context) {
	var req struct {
		Published bool `json:"published"`
	}
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.PublishProgram(paramID(ctx), req.Published, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	if value == nil {
		response.Error(ctx, http.StatusNotFound, response.NotFound, "Program kerja tidak ditemukan")
		return
	}
	response.OK(ctx, "Status program kerja diperbarui", value)
}
func (c *Controller) CreateMilestone(ctx *gin.Context) {
	var req dto.MilestoneRequest
	if !bind(ctx, &req) {
		return
	}
	req.WorkProgramID = paramID(ctx)
	value, err := c.service.CreateMilestone(req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Milestone dibuat", value))
}
func (c *Controller) CreateDocumentation(ctx *gin.Context) {
	var req dto.DocumentationRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.CreateDocumentation(paramID(ctx), req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Dokumentasi ditambahkan", value))
}
func (c *Controller) ReorderDocumentation(ctx *gin.Context) {
	var req struct {
		IDs []uint64 `json:"ids"`
	}
	if !bind(ctx, &req) {
		return
	}
	if err := c.service.ReorderDocumentation(paramID(ctx), req.IDs, middlewares.CurrentClaims(ctx)); err != nil {
		response.Error(ctx, http.StatusForbidden, response.Forbidden, err.Error())
		return
	}
	response.OK(ctx, "Urutan dokumentasi diperbarui", nil)
}
func (c *Controller) CreateMedia(ctx *gin.Context) {
	var req dto.MediaRequest
	if !bind(ctx, &req) {
		return
	}
	value, err := c.service.CreateMedia(req, middlewares.CurrentClaims(ctx))
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response.BuildResponseSuccess("Media asset dibuat", value))
}

func bind(ctx *gin.Context, value any) bool {
	if err := ctx.ShouldBindJSON(value); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, response.ValidationError, "Payload tidak valid")
		return false
	}
	return true
}
func paramID(ctx *gin.Context) uint64 { id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64); return id }
func queryInt(ctx *gin.Context, key string, fallback int) int {
	value, err := strconv.Atoi(ctx.Query(key))
	if err != nil || value < 1 {
		return fallback
	}
	return value
}
func writePage(ctx *gin.Context, message string, page, perPage int, total int64, totalPages int, data any) {
	response.List(ctx, message, data, response.Meta{Page: page, PerPage: perPage, Total: total, TotalPages: totalPages})
}

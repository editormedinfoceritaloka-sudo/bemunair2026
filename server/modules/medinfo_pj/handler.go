package medinfo_pj

import (
	"net/http"
	"strconv"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct{ repo *Repository }

func NewHandler(repo *Repository) *Handler { return &Handler{repo: repo} }

func (h *Handler) List(c *gin.Context) {
	rows, err := h.repo.List()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Gagal mengambil queue")
		return
	}
	response.List(c, "Queue PJ Medinfo", rows, response.Meta{Page: 1, PerPage: len(rows), Total: int64(len(rows)), TotalPages: 1})
}

func (h *Handler) Create(c *gin.Context) {
	var req struct {
		UserID   uint64 `json:"user_id"`
		Position int    `json:"position"`
	}
	if c.ShouldBindJSON(&req) != nil || req.UserID == 0 {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	row := &entities.MedinfoPJQueue{UserID: req.UserID, Position: req.Position}
	if row.Position == 0 {
		row.Position = 1
	}
	if err := h.repo.Create(row); err != nil {
		response.Error(c, http.StatusConflict, response.Conflict, "Queue gagal dibuat")
		return
	}
	response.Created(c, "Queue berhasil dibuat", row)
}

func (h *Handler) Reorder(c *gin.Context) {
	var req struct {
		IDs []uint64 `json:"ids"`
	}
	if c.ShouldBindJSON(&req) != nil || len(req.IDs) == 0 {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	if err := h.repo.Reorder(req.IDs); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Queue gagal diurutkan")
		return
	}
	response.OK(c, "Queue berhasil diurutkan", nil)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.repo.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Queue gagal dihapus")
		return
	}
	response.OK(c, "Queue berhasil dihapus", nil)
}

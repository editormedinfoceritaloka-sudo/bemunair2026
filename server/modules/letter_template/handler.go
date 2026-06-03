package letter_template

import (
	"bemunair2026/server/database/entities"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct{ repo *Repository }

func NewHandler(repo *Repository) *Handler { return &Handler{repo: repo} }
func (h *Handler) Create(c *gin.Context) {
	var t entities.LetterTemplate
	if c.ShouldBindJSON(&t) != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	if err := h.repo.Create(&t); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Template gagal dibuat")
		return
	}
	response.Created(c, "Template berhasil dibuat", t)
}
func (h *Handler) List(c *gin.Context) {
	rows, err := h.repo.List()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Gagal mengambil template")
		return
	}
	response.List(c, "Daftar template", rows, response.Meta{Page: 1, PerPage: len(rows), Total: int64(len(rows)), TotalPages: 1})
}
func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	row, err := h.repo.FindByID(id)
	if err != nil || row == nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "Template tidak ditemukan")
		return
	}
	response.OK(c, "Detail template", row)
}
func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	row, err := h.repo.FindByID(id)
	if err != nil || row == nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "Template tidak ditemukan")
		return
	}
	if c.ShouldBindJSON(row) != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	row.ID = id
	if err := h.repo.Update(row); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Template gagal diperbarui")
		return
	}
	response.OK(c, "Template berhasil diperbarui", row)
}
func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.repo.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Template gagal dihapus")
		return
	}
	response.OK(c, "Template berhasil dihapus", nil)
}

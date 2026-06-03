package user

import (
	"net/http"
	"strconv"

	"bemunair2026/server/database/entities"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct{ repo *Repository }

func NewHandler(repo *Repository) *Handler { return &Handler{repo: repo} }

func (h *Handler) List(c *gin.Context) {
	users, err := h.repo.List()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Gagal mengambil user")
		return
	}
	response.List(c, "Daftar user", users, response.Meta{Page: 1, PerPage: len(users), Total: int64(len(users)), TotalPages: 1})
}

func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	u, err := h.repo.FindByID(id)
	if err != nil || u == nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "User tidak ditemukan")
		return
	}
	response.OK(c, "Detail user", u)
}

func (h *Handler) Create(c *gin.Context) {
	var req struct {
		Name, Email, Password, Role string
		Ministry, Phone             *string
	}
	if c.ShouldBindJSON(&req) != nil || req.Email == "" || req.Password == "" {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	u := &entities.User{Name: req.Name, Email: req.Email, PasswordHash: string(hash), Role: req.Role, Ministry: req.Ministry, Phone: req.Phone}
	if err := h.repo.Create(u); err != nil {
		response.Error(c, http.StatusConflict, response.Conflict, "User gagal dibuat")
		return
	}
	response.Created(c, "User berhasil dibuat", u)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	u, err := h.repo.FindByID(id)
	if err != nil || u == nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "User tidak ditemukan")
		return
	}
	var req struct {
		Name, Email, Role string
		Ministry, Phone   *string
	}
	if c.ShouldBindJSON(&req) != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	if req.Name != "" {
		u.Name = req.Name
	}
	if req.Email != "" {
		u.Email = req.Email
	}
	if req.Role != "" {
		u.Role = req.Role
	}
	u.Ministry = req.Ministry
	u.Phone = req.Phone
	if err := h.repo.Update(u); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "User gagal diperbarui")
		return
	}
	response.OK(c, "User berhasil diperbarui", u)
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.repo.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "User gagal dihapus")
		return
	}
	response.OK(c, "User berhasil dihapus", nil)
}

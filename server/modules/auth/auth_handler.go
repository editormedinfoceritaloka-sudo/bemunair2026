package auth

import (
	"net/http"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/user"
	"bemunair2026/server/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
	users   *user.Repository
}

func NewHandler(service *Service, users *user.Repository) *Handler {
	return &Handler{service: service, users: users}
}

type registerRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Role     string  `json:"role"`
	Ministry *string `json:"ministry"`
	Phone    *string `json:"phone"`
}

func (h *Handler) Register(c *gin.Context) {
	var req registerRequest
	if c.ShouldBindJSON(&req) != nil || req.Email == "" || req.Password == "" {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	u, err := h.service.Register(req.Name, req.Email, req.Password, req.Role, req.Ministry, req.Phone)
	if err != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, err.Error())
		return
	}
	response.Created(c, "User berhasil dibuat", u)
}

func (h *Handler) Login(c *gin.Context) {
	var req struct{ Email, Password string }
	if c.ShouldBindJSON(&req) != nil || req.Email == "" || req.Password == "" {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	token, u, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, response.Unauthenticated, "Kredensial salah")
		return
	}
	response.OK(c, "Login berhasil", gin.H{"token": token, "user": u})
}

func (h *Handler) Me(c *gin.Context) {
	claims := middlewares.CurrentClaims(c)
	u, err := h.users.FindByID(claims.UserID)
	if err != nil || u == nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "User tidak ditemukan")
		return
	}
	response.OK(c, "User aktif", u)
}

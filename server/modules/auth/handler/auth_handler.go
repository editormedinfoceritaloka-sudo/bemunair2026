package handler

import (
	"net/http"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/auth/dto"
	"bemunair2026/server/modules/auth/service"
	"bemunair2026/server/modules/auth/validation"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service    *service.Service
	validation *validation.AuthValidation
}

func NewHandler(service *service.Service, validation *validation.AuthValidation) *Handler {
	return &Handler{service: service, validation: validation}
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if c.ShouldBindJSON(&req) != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	if err := h.validation.ValidateRegisterRequest(req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	user, err := h.service.Register(req)
	if err != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, err.Error())
		return
	}
	response.Created(c, "User berhasil dibuat", user)
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if c.ShouldBindJSON(&req) != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	if err := h.validation.ValidateLoginRequest(req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, response.ValidationError, "Validasi gagal")
		return
	}
	res, err := h.service.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, response.Unauthenticated, "Kredensial salah")
		return
	}
	response.OK(c, "Login berhasil", res)
}

func (h *Handler) Me(c *gin.Context) {
	claims := middlewares.CurrentClaims(c)
	user, err := h.service.Me(claims.UserID)
	if err != nil || user == nil {
		response.Error(c, http.StatusNotFound, response.NotFound, "User tidak ditemukan")
		return
	}
	response.OK(c, "User aktif", user)
}

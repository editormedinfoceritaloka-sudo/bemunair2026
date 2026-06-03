package controller

import (
	"net/http"

	"bemunair2026/server/middlewares"
	"bemunair2026/server/modules/auth/dto"
	"bemunair2026/server/modules/auth/service"
	"bemunair2026/server/modules/auth/validation"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Me(ctx *gin.Context)
}

type authController struct {
	authService    service.AuthService
	authValidation *validation.AuthValidation
}

var _ AuthController = (*authController)(nil)

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService:    authService,
		authValidation: validation.NewAuthValidation(),
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}
	if err := c.authValidation.ValidateRegisterRequest(req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := c.authService.Register(req)
	if err != nil {
		res := response.BuildResponseFailed("User gagal dibuat", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := response.BuildResponseSuccess("User berhasil dibuat", user)
	ctx.JSON(http.StatusCreated, res)
}

func (c *authController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}
	if err := c.authValidation.ValidateLoginRequest(req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	result, err := c.authService.Login(req)
	if err != nil {
		res := response.BuildResponseFailed("Kredensial salah", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	res := response.BuildResponseSuccess("Login berhasil", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) Me(ctx *gin.Context) {
	claims := middlewares.CurrentClaims(ctx)
	user, err := c.authService.Me(claims.UserID)
	if err != nil || user == nil {
		res := response.BuildResponseFailed("User tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("User aktif", user)
	ctx.JSON(http.StatusOK, res)
}

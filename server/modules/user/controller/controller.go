package controller

import (
	"net/http"
	"strconv"

	"bemunair2026/server/modules/user/dto"
	"bemunair2026/server/modules/user/service"
	"bemunair2026/server/modules/user/validation"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userController struct {
	userService    service.UserService
	userValidation *validation.UserValidation
}

var _ UserController = (*userController)(nil)

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService:    userService,
		userValidation: validation.NewUserValidation(),
	}
}

func (c *userController) List(ctx *gin.Context) {
	users, err := c.userService.List()
	if err != nil {
		res := response.BuildResponseFailed("Gagal mengambil user", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("Daftar user", users)
	res.Meta = response.Meta{Page: 1, PerPage: len(users), Total: int64(len(users)), TotalPages: 1}
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Get(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	user, err := c.userService.Get(id)
	if err != nil || user == nil {
		res := response.BuildResponseFailed("User tidak ditemukan", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("Detail user", user)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Create(ctx *gin.Context) {
	var req dto.UserCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}
	if err := c.userValidation.ValidateCreateRequest(req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := c.userService.Create(req)
	if err != nil {
		res := response.BuildResponseFailed("User gagal dibuat", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}

	res := response.BuildResponseSuccess("User berhasil dibuat", user)
	ctx.JSON(http.StatusCreated, res)
}

func (c *userController) Update(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var req dto.UserUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}
	if err := c.userValidation.ValidateUpdateRequest(req); err != nil {
		res := response.BuildResponseFailed("Validasi gagal", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := c.userService.Update(id, req)
	if err != nil || user == nil {
		res := response.BuildResponseFailed("User gagal diperbarui", response.NotFound, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := response.BuildResponseSuccess("User berhasil diperbarui", user)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.userService.Delete(id); err != nil {
		res := response.BuildResponseFailed("User gagal dihapus", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess("User berhasil dihapus", nil)
	ctx.JSON(http.StatusOK, res)
}

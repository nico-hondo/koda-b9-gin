package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nico-hondo/internal/dto"
)

type IUserService interface {
	EmptyValidationUser(data dto.User) error
	AvailableUser(data dto.User) error
	Authentication(user dto.LoginInput) error
}

type UserHandler struct {
	pu IUserService
}

func NewUserHandler(pu IUserService) *UserHandler {
	return &UserHandler{
		pu: pu,
	}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	var data dto.User

	if err := ctx.ShouldBindWith(&data, binding.JSON); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     "terjadi kesalahan server",
		})
		return
	}

	if err := u.pu.EmptyValidationUser(data); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     err.Error(),
		})
		return
	}

	if err := u.pu.AvailableUser(data); err != nil {
		ctx.JSON(http.StatusForbidden, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     err.Error(),
		})
		return
	}

	dto.Users = append(dto.Users, data)

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Data:    dto.Users,
		Msg:     "Register Berhasil",
	})
}

func (u *UserHandler) LoginUser(ctx *gin.Context) {
	var login dto.LoginInput

	if err := ctx.ShouldBindWith(&login, binding.FormPost); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     err.Error(),
		})
	}

	if err := u.pu.Authentication(login); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    dto.Users,
		Msg:     "Login Berhasil",
	})
}

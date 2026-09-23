package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nico-hondo/internal/dto"
)

type IPingService interface {
	EmptyValidation(body dto.Body) error
}

type PingHandler struct {
	ps IPingService
}

func NewPingHandler(ps IPingService) *PingHandler {
	return &PingHandler{
		ps: ps,
	}
}

// Deklarasi router
func (p *PingHandler) Pong(ctx *gin.Context) {
	//Send Response
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Pong",
	})
}

func (p *PingHandler) Greet(ctx *gin.Context) {

	var body dto.Body
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		// log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     "terjadi kesalahan server",
		})

		return
	}

	//Business Logic, Tarik Data
	if err := p.ps.EmptyValidation(body); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     "terjadi kesalahan",
		})
		return
	}

	//success response
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    body,
		Msg:     fmt.Sprintf("Selamat Datang %s", body.Name),
	})
}

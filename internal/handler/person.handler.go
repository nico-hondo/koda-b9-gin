package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/service"
)

type PersonHandler struct {
	ps *service.PersonService
}

func NewPersonHandler(ps *service.PersonService) *PersonHandler {
	return &PersonHandler{
		ps: ps,
	}
}

func (p *PersonHandler) GetAllPerson(c *gin.Context) {
	persons, err := p.ps.GetAllPerson(c.Request.Context())
	if err != nil {
		log.Println("error: ", err.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    persons,
		Msg:     "Data Person berhasil diambil",
	})
}

package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/nico-hondo/internal/dto"
)

func initHeaderRouter(r *gin.Engine) {
	headerRouter := r.Group("/header")

	// headerRouter.GET("", func(ctx *gin.Context) {

	// })

	headerRouter.GET("event/:id/:slug", func(ctx *gin.Context) {
		id := ctx.Param("id")
		slug := ctx.Param("slug")

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"id":   id,
				"slug": slug,
			},
		})
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", func(fl validator.FieldLevel) bool {
			date, ok := fl.Field().Interface().(time.Time)
			if !ok {
				return false
			}
			// contoh: gak boleh booking tanggal lewat / di masa lalu
			today := time.Now().Truncate(24 * time.Hour)
			return date.After(today) || date.Equal(today)
		})
	}

	headerRouter.GET("booking", func(ctx *gin.Context) {
		var b dto.Booking
		if err := ctx.ShouldBindQuery(&b); err != nil {
			fmt.Println("BIND ERROR:", err)
			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Data:    nil,
				Msg:     err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"check_in":  b.CheckIn.Format("2006-01-02"),
				"check_out": b.CheckOut.Format("2006-01-02"),
			},
			Msg: "Booking dates are Valid",
		})
	})

	headerRouter.GET("query", func(ctx *gin.Context) {
		title := ctx.Query("title")
		genre := ctx.QueryArray("genre")

		var m dto.Movie
		if err := ctx.ShouldBindWith(&m, binding.Query); err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"title": title,
				"genre": genre,
			},
			Msg: "Booking dates are Valid",
		})
	})
}

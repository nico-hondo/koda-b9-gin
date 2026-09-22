package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/router"
)

// type Response map[string]any

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		//Send Response
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "World",
		})
	})

	r.GET("/how", func(ctx *gin.Context) {
		//Send Response
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "good",
		})
	})

	router.InitMainRouter(r)

	r.Run("localhost:8000")
}

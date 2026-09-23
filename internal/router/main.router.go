package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/middleware"
)

func InitMainRouter(r *gin.Engine) {
	r.Use(middleware.Cors)

	initPingRouter(r)
	initUserRouter(r)
	initHeaderRouter(r)
}

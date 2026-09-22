package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/service"
)

func initPingRouter(r *gin.Engine) {
	//Deklarasi router
	pingRouter := r.Group("/ping")

	ps := service.NewPingService()

	ph := handler.NewPingHandler(ps)

	pingRouter.GET("", ph.Pong)
	pingRouter.POST("", ph.Greet)
}

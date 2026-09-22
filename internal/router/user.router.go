package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/service"
)

func initUserRouter(r *gin.Engine) {
	userRouter := r.Group("/auth")

	ps := service.NewUserService()

	ph := handler.NewUserHandler(ps)

	userRouter.POST("", ph.LoginUser)
	userRouter.POST("/register", ph.Register)
}

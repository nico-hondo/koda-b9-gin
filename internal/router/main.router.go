package router

import "github.com/gin-gonic/gin"

func InitMainRouter(r *gin.Engine) {
	initPingRouter(r)
	initUserRouter(r)
}

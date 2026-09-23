package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	allowedOrigin := []string{"http://127.0.0.1:5500"}
	currentOrigin := ctx.GetHeader("Origin")
	if slices.Contains(allowedOrigin, currentOrigin) {
		ctx.Header("Access-Control-Allow-Origin", currentOrigin)
	}

	ctx.Header("Access-Control-Allow-Headers", "Content-Type, XXX-Header")
	ctx.Header("Access-Control-Allow-Method", "GET, OPTIONS, PATCH")

	// preflight cors
	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}

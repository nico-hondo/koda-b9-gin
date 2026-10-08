package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/middleware"
)

func InitMainRouter(r *gin.Engine, db *pgxpool.Pool) {
	r.Use(middleware.Cors)

	initPingRouter(r)
	initUserRouter(r)
	initHeaderRouter(r)
	initPersonRouter(r, db)
}

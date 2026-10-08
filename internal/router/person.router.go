package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/internal/service"
)

func initPersonRouter(r *gin.Engine, db *pgxpool.Pool) {
	personRouter := r.Group("/person")

	pr := repo.NewPersonRepo(db)
	ps := service.NewPersonService(pr)
	ph := handler.NewPersonHandler(ps)

	personRouter.GET("", ph.GetAllPerson)

}

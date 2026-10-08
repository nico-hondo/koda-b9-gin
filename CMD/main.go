package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nico-hondo/internal/config"
	"github.com/nico-hondo/internal/router"
)

// type Response map[string]any

func main() {

	//Load env seawal mungkin
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect ke DB
	pdb := config.NewPsqlDb(os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))

	pool, err := pdb.Connect()

	if err != nil {
		log.Println("Cannot connect to db\nReason: ", err.Error())
		return
	}

	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Println("Database is not ready\nReason: ", err.Error())
		return
	}

	log.Println("Database Ready")

	// Generate gin engine
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

	// Deklarasi Router (endpoint & method HTTP)
	router.InitMainRouter(r, pool)

	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}

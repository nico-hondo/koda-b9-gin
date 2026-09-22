package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// type Response map[string]any

type Body struct {
	Name string `form:"nama" json:"nama"`
	Age  int8   `form:"umur" json:"umur"`
}

type Response struct {
	Success bool
	Data    any
	Msg     string
}

type User struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

var Users []User

type LoginInput struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func main() {
	router := gin.Default()

	//Deklarasi router
	router.GET("/ping", func(ctx *gin.Context) {
		//Send Response
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Pong",
		})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		//Send Response
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "World",
		})
	})

	router.GET("/how", func(ctx *gin.Context) {
		//Send Response
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "good",
		})
	})

	router.POST("/add", func(ctx *gin.Context) {
		var body Body
		if err := ctx.ShouldBindWith(&body, binding.Form); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})

			return
		}

		if body.Name == "" && body.Age == 0 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan",
			})
			return
		}

		//misal tidak nil
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Data:    body,
			Msg:     fmt.Sprintf("Selamat Datang %s", body.Name),
		})
	})

	router.POST("/register", func(ctx *gin.Context) {
		var data User
		if err := ctx.ShouldBindWith(&data, binding.FormPost); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		if data.Email == "" && data.Password == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan",
			})
			return
		}

		for _, val := range Users {
			if data.Email == val.Email {
				ctx.JSON(http.StatusForbidden, Response{
					Success: false,
					Data:    nil,
					Msg:     fmt.Sprintf("email %s sudah tersedia", data.Email),
				})
				return
			}
		}

		Users = append(Users, data)

		ctx.JSON(http.StatusCreated, Response{
			Success: true,
			Data:    Users,
			Msg:     "Register Berhasil",
		})
	})

	router.POST("/login", func(ctx *gin.Context) {
		var login LoginInput

		if err := ctx.ShouldBindWith(&login, binding.FormPost); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
		}

		for _, val := range Users {
			if login.Email != val.Email || login.Password != val.Password {
				ctx.JSON(http.StatusBadRequest, Response{
					Success: false,
					Data:    nil,
					Msg:     "nama atau password salah",
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Data:    Users,
			Msg:     "Login Berhasil",
		})
	})

	router.Run("localhost:8000")
}

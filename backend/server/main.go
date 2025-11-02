package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/golang_fullstack_ai_movie_streaming_app/backend/server/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/movie", controller.AddMovie())
	router.POST("/user/register", controller.RegisterUser())


	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server",err)
	}
}
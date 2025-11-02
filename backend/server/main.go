package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang_fullstack_ai_movie_streaming_app/backend/server/routes"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})

	routes.SetupProtectedRoutes(router)
	routes.SetupUnProtectedRouted(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}

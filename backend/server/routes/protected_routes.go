package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/golang_fullstack_ai_movie_streaming_app/backend/server/controllers"
	"github.com/golang_fullstack_ai_movie_streaming_app/backend/server/middlerware"
)

func SetupProtectedRoutes(router *gin.Engine) {


	protected := router.Group("/")
	protected.Use(middlerware.AuthMiddleWare())
	protected.GET("/movie/:imdb_id", controller.GetMovie())
	protected.POST("/movie", controller.AddMovie())
}
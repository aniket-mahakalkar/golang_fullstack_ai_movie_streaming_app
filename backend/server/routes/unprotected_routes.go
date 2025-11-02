package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/golang_fullstack_ai_movie_streaming_app/backend/server/controllers"
)

func SetupUnProtectedRouted(router *gin.Engine) {

	router.GET("/movies", controller.GetMovies())
	router.POST("/user/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())

}

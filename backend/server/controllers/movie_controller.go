package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang_fullstack_ai_movie_streaming_app/backend/server/database"
	"github.com/golang_fullstack_ai_movie_streaming_app/backend/server/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

func GetMovies() gin.HandlerFunc{

	return func (c *gin.Context)  {

		ctx , cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		
		var movies []models.Movie

		cursor, err :=  movieCollection.Find(ctx, bson.M{})

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		
		
		}

		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, movies)
	
	}
}


func GetMovie() gin.HandlerFunc {

	return  func(c *gin.Context) {

		ctx, cancle := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancle()

		movieID := c.Param("imdb_id")

		if movieID == "" {

			c.JSON(http.StatusBadRequest, gin.H{"error":"Movie ID is required"})
			return 
		}

		var movie models.Movie

		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)

		if err != nil {

			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return 
		
		}


		c.JSON(http.StatusOK, movie)
		

	}
}
package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang_fullstack_ai_movie_streaming_app/backend/server/database"
	"github.com/golang_fullstack_ai_movie_streaming_app/backend/server/models"
	"github.com/golang_fullstack_ai_movie_streaming_app/backend/server/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)



var userCollection *mongo.Collection = database.OpenCollection("users")

func HashPassword(password string) (string, error){

	HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if  err != nil  {

		return  "", err
		
	}

	return  string(HashPassword) , nil
}

func RegisterUser() gin.HandlerFunc{

	return func(c *gin.Context) {

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Inpute data"})
			return 
		}

		validate := validator.New()

		if err := validate.Struct(user); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error":"Inavlid request body", "details":err.Error()})
			return 
		}

		hashedPassword, err := HashPassword(user.Password)

		if err != nil {

			c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to hash the password"})
		}

		var ctx, cancle = context.WithTimeout(context.Background(),100 * time.Second)

		defer cancle()

		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to check existing user"})
			return 
		}

		if count > 0 {

			c.JSON(http.StatusConflict, gin.H{"error":"User already exists"})
			return 
		}

		user.UserID = bson.NewObjectID().Hex()
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.Password = hashedPassword

		result, err := userCollection.InsertOne(ctx, user)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create user"})
			return 
		}

		c.JSON(http.StatusCreated, result)
		
	}
}

func LoginUser() gin.HandlerFunc{

	return func(c *gin.Context) {

		var UserLogin models.UserLogin

		if err := c.ShouldBindJSON(&UserLogin); err != nil{

			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input data"})
			return 
		}

		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancle()

		var foundUser models.User

		err := userCollection.FindOne(ctx, bson.M{"email":UserLogin.Email}).Decode(&foundUser)

		if err != nil {

			c.JSON(http.StatusNotFound, gin.H{"error":"User not found"})
			return 
		}

		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(UserLogin.Password))

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid credentials"})
			return 
		}

		token, refreshToken, err := utils.GenerateAllTokens(foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.Role,foundUser.UserID)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to generate tokens"})
		}

		err = utils.UpdateAllTokens(foundUser.UserID, token, refreshToken)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to update token"})

			return 
		}

		c.JSON(http.StatusOK, models.UserResponse{
			UserId: foundUser.UserID,
			FirstName: foundUser.FirstName,
			LastName: foundUser.LastName,
			Email: foundUser.Email,
			Role: foundUser.Role,
			Token: token,
			RefreshToken: refreshToken,
			FavouriteGenres: foundUser.FavouriteGenres,
		})
	}
}
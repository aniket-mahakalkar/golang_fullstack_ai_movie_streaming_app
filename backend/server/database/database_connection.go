package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {

	err := godotenv.Load(".env") 

	if err != nil {
		log.Println("Error loading .env file")
	
	}

	MongoDb := os.Getenv("MONGODB_URI")

	if MongoDb == ""{
		log.Fatal("MONGODB_URI is not set!")
	}

	fmt.Println("MongoDB URI: ",MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)

	client , err := mongo.Connect(clientOptions)
	

	if err != nil {

		log.Println("Error occured while connection to mongo db")
	}

	return  client

}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")


	fmt.Println("Database Name: ",databaseName)

	collection :=  Client.Database(databaseName).Collection(collectionName)

	if collection == nil {

		return  nil
	}

	return collection
	
}
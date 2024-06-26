package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ajaypp123/golang-jwt-microservice/helpers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Not found .env value loading from env MONGODB_URL")
	}
	MongoDb := "mongodb://" + helpers.GetConfig().Mongo.Server + ":27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("go-auth").Collection(collectionName)
	return collection
}

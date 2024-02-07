package db

import (
	"context"
	"log"
	"os"

	"gopad/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return *client
}

func CreatePerson(person *models.Person) {
	client := GetMongoClient()
	_, err := client.Database("mcrap_db").Collection("person").InsertOne(context.TODO(), &person)
	if err != nil {
		log.Fatal(err)
	}
}

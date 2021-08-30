package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var onceDatabase sync.Once
var onceClient sync.Once

var database *mongo.Database
var mongoClient *mongo.Client

func getMongoClient() *mongo.Client {
	onceClient.Do(func() {
		var err error
		mongoUri := os.Getenv("DB_CNN")
		mongoClient, err = mongo.NewClient(options.Client().ApplyURI(mongoUri))
		if err != nil {
			log.Panic(err)
		}
	})

	return mongoClient
}

func GetMongoDatabase() *mongo.Database {
	onceDatabase.Do(func() {
		mongoClient := getMongoClient()
		databaseName := os.Getenv("DATABASE_NAME")
		database = mongoClient.Database(databaseName)
	})

	return database
}

func IsConnectedToMongoDB() bool {
	client := getMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := client.Connect(ctx)

	return err == nil
}

func CloseConnection() {
	client := getMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client.Disconnect(ctx)
}

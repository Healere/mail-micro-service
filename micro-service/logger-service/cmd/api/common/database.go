package common

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() {
	// connect to mongo
	mongoClient, err := ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient
	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func ConnectToMongo() (*mongo.Client, error) {
	mongoURL := viper.GetString("mongo.mongoURL")
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")
	// Verify if the url is valid
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connect to mongodb: ", err)
		return nil, err
	}

	return c, nil
}

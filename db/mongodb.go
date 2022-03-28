package db

import (
	"context"
	"log"
	"my_gym_go/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func InitMongoDb(config *config.Configuration) *mongo.Client {

	c, err := mongo.NewClient(options.Client().ApplyURI(config.URLConexaoDB))
	if err != nil {
		log.Panicln(err)
	}

	c.Connect(context.Background())
	client = c

	return c
}

func GetDatabase(name string) *mongo.Database {
	return client.Database(name)
}

package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:pass12345@localhost:27017/?maxPoolSize=20&w=majority"

func InitMongoDb() (*mongo.Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		defer cancel()
		return nil, err
	}

	defer cancel()
	defer client.Disconnect(ctx)

	return client, nil
}

package utils

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbClient struct {
	client *mongo.Client
	ctx    context.Context
}

var (
	mongodb_client_instance *MongoDbClient
	once_mongodb_client     sync.Once
)

func GetMongoDbClient() *MongoDbClient {
	once_mongodb_client.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}

		// Verify the client is connected
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		mongodb_client_instance = &MongoDbClient{client: client, ctx: ctx}
	})
	return mongodb_client_instance
}

func (mc *MongoDbClient) GetClient() *mongo.Client {
	return mc.client
}

func (mc *MongoDbClient) Disconnect() {
	if mc.client != nil {
		mc.client.Disconnect(mc.ctx)
		log.Println("DB disconnection was successful")
	}
}

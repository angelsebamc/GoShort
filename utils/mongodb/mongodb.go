package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	client *mongo.Client
	ctx    context.Context
}

var mongodb_instance *MongoDb

func GetInstance() *MongoDb {
	if mongodb_instance == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION")))
		if err != nil {
			log.Fatal(err)
		}

		// Verify the client is connected
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		mongodb_instance = &MongoDb{client: client, ctx: ctx}
		log.Println("Connected to MongoDB")
	}

	return mongodb_instance
}

func (m *MongoDb) Disconnect() {
	if mongodb_instance.client != nil {
		mongodb_instance.client.Disconnect(m.ctx)
		log.Println("DB disconnection was successful")
	}
}

// getters
func (m *MongoDb) GetClient() *mongo.Client {
	return m.client
}

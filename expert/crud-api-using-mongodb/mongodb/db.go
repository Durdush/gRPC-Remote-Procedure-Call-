package mongodb

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

const (
	DB_NAME        = "Informatica"
	DB_MONGODB_URL = "mongodb://localhost:27017"
)

func InitDB() error {
	client, err := mongo.NewClient(DB_MONGODB_URL)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return err
}

func CreateCollection(collectionName string) *mongo.Collection {
	return client.Database(DB_NAME).Collection(collectionName)
}

func CloseMongoDB() {
	client.Disconnect(context.TODO())
}

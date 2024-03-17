package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client

func Init(ctx context.Context) error {
	dbHost := os.Getenv("MONGODB_HOST")
	dbPort := os.Getenv("MONGODB_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	connOption := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, connOption)
	if err != nil {
		return err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}
	log.Print("Ping mongodb success")
	Client = client
	return nil
}
func GetCollection(collection string) *mongo.Collection {
	database := os.Getenv("MONGODB_DATABASE")
	return Client.Database(database).Collection(collection)
}

func Disconnect(ctx context.Context) {
	if err := Client.Disconnect(ctx); err != nil {
		log.Fatal("Disconnect mongodb fail ", err)
	}
	log.Print("Disconnect mongodb success")

}

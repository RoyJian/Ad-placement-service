package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client *mongo.Client

func Init(ctx context.Context) error {
	dbHost := os.Getenv("MONGODB_HOST")
	dbPort := os.Getenv("MONGODB_PORT")
	cred := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    os.Getenv("MONGODB_DATABASE"),
		Username:      os.Getenv("MONGODB_ADMIN"),
		Password:      os.Getenv("MONGODB_PASSWORD"),
	}
	uri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	connOption := options.Client().ApplyURI(uri).SetAuth(cred)
	c, err := mongo.Connect(ctx, connOption)
	if err != nil {
		return err
	}
	if err := c.Ping(ctx, nil); err != nil {
		return err
	}
	log.Print("Ping mongodb success")
	client = c
	return nil
}
func GetCollection(collection string) *mongo.Collection {
	database := os.Getenv("MONGODB_DATABASE")
	return client.Database(database).Collection(collection)
}

func Disconnect(ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal("Disconnect mongodb fail ", err)
	}
	log.Print("Disconnect mongodb success")

}

package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"log"
	"os"
)

type Db interface {
	Disconnect(ctx context.Context)
	GetCollection(collection string) *mongo.Collection
}

type MongoDb struct {
	client *mongo.Client
}

func NewMongoDB() *MongoDb {
	ctx := context.Background()
	dbHost := os.Getenv("MONGODB_HOST")
	dbPort := os.Getenv("MONGODB_PORT")
	cred := options.Credential{
		//AuthMechanism: "SCRAM-SHA-1",
		//AuthSource:    os.Getenv("MONGODB_DATABASE"),
		Username: os.Getenv("MONGODB_ADMIN"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	}
	uri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	connOption := options.Client().ApplyURI(uri).SetAuth(cred)
	client, err := mongo.Connect(ctx, connOption)
	if err != nil {
		log.Fatal("Connect MongoDB error", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Ping mongodb error", err)
	}
	log.Println("Ping mongodb success")
	return &MongoDb{client: client}
}

func (m MongoDb) GetCollection(collection string) *mongo.Collection {
	database := os.Getenv("MONGODB_DATABASE")
	return m.client.Database(database).Collection(collection)
}

func (m MongoDb) Disconnect(ctx context.Context) {
	if err := m.client.Disconnect(ctx); err != nil {
		log.Fatal("Disconnect mongodb fail ", err)
	}
	log.Print("Disconnect mongodb success")
}

var Module = fx.Options(
	fx.Provide(NewMongoDB))

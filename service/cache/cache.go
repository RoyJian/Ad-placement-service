package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

var client *redis.ClusterClient

func Init(ctx context.Context) error {
	address := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	opt := &redis.ClusterOptions{
		Addrs:    []string{address},
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	c := redis.NewClusterClient(opt)
	client = c
	err := client.ForEachMaster(ctx, func(ctx context.Context, shard *redis.Client) error {
		log.Println("ping ", shard.String())
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	log.Println("Ping redis success")
	return nil
}

func Get(ctx context.Context, key string) *redis.StringCmd {
	data := client.Get(ctx, key)
	return data
}

func Set(key string, value interface{}) error {
	ctx := context.Background()
	if err := client.Set(ctx, key, value, 3*time.Hour).Err(); err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

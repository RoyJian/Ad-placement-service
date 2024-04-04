package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
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

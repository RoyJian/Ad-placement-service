package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"log"
	"os"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, ket string, value interface{}) error
}

type Redis struct {
	client *redis.ClusterClient
}

func NewRedis() *Redis {
	ctx := context.Background()
	address := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	opt := &redis.ClusterOptions{
		Addrs:    []string{address},
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	client := redis.NewClusterClient(opt)
	err := client.ForEachMaster(ctx, func(ctx context.Context, shard *redis.Client) error {
		log.Println("ping ", shard.String())
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		log.Fatal("Ping Redis  Error", err)
	}
	log.Println("Ping redis success")
	return &Redis{client: client}
}

func (r Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	data := r.client.Get(ctx, key)
	return data
}

func (r Redis) Set(ctx context.Context, key string, value interface{}) error {
	if err := r.client.Set(ctx, key, value, 3*time.Hour).Err(); err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

var Module = fx.Options(
	fx.Provide(NewRedis))

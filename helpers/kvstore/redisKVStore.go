package kvstore

import (
	"context"
	"strconv"

	"github.com/ajaypp123/golang-jwt-microservice/helpers"
	"github.com/go-redis/redis/v8"
)

type KVStoreI interface {
	Set(ctx context.Context, key string, val *string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type KVstore struct {
	client *redis.Client
	db     int
}

func GetKVStore() KVStoreI {
	return &KVstore{
		client: redis.NewClient(&redis.Options{
			Addr:     helpers.GetConfig().Redis.Server + ":" + strconv.Itoa(helpers.GetConfig().Redis.Port),
			Password: helpers.GetConfig().Redis.Password,
			DB:       helpers.GetConfig().Redis.Database,
		}),
		db: helpers.GetConfig().Redis.Database,
	}
}

// Del implements KVStoreI.
func (k *KVstore) Del(ctx context.Context, key string) error {
	return k.client.Del(ctx, "key").Err()
}

// Get implements KVStoreI.
func (k *KVstore) Get(ctx context.Context, key string) (string, error) {
	return k.client.Get(ctx, key).Result()
}

// Set implements KVStoreI.
func (k *KVstore) Set(ctx context.Context, key string, val *string) error {
	return k.client.Set(ctx, key, &val, 0).Err()
}

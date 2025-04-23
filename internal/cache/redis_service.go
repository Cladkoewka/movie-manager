package cache

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Cladkoewka/movie-manager/internal/model/dto"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService() *RedisService {
	client := newRedisClient()

	return &RedisService{client: client}
}

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", 
		DB:       0,  
	})

	return client
}

func (r *RedisService) GetCache(ctx context.Context, key string, dest interface{}) error {
	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

func (r *RedisService) SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}	

	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *RedisService) GenerateCacheKey(prefix string, params dto.MovieQueryParams) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	hash := sha1.Sum(bytes) 
	return fmt.Sprintf("%s:%s", prefix, hex.EncodeToString(hash[:])), nil
}
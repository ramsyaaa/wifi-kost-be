package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

// New creates a new go-redis client.
func New(addr string, passwd string, db int) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})

	return &Cache{client}
}

// Ping checks whether connection is possible to Redis.
func (c *Cache) Ping(ctx context.Context) error {
	return c.Client.Ping(ctx).Err()
}

// Set stores key with a value in Redis.
func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.Client.Set(ctx, key, value, ttl).Err()
}

// Get retrieves a value from Redis.
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}

package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
  *redis.Client
}

func New(addr string, passwd string, db  int) *Cache{
 Client := redis.NewClient(&redis.Options{
	Addr:     addr,
	Password: passwd,
	DB:       db,
 })

 return &Cache{Client}
}

func (c *Cache) Ping(ctx context.Context) error {
	 return c.Client.Ping(ctx).Err()
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.Client.Set(ctx, key, value, ttl).Err()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}
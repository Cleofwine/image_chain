package cache

import (
	"context"
	redis_pkg "shorturl-grpc/pkg/db/redis"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisKVCache struct {
	redisClient *redis.Client
	destroy     func()
}

func newRedisKVCache(client *redis.Client, destory func()) KVCache {
	return &redisKVCache{
		redisClient: client,
		destroy:     destory,
	}
}

func (c *redisKVCache) Get(key string) (string, error) {
	key = redis_pkg.GetKey(key)
	res, err := c.redisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return res, err
}
func (c *redisKVCache) Set(key, value string, ttl int) error {
	key = redis_pkg.GetKey(key)
	return c.redisClient.SetEx(context.Background(), key, value, time.Second*time.Duration(ttl)).Err()
}
func (c *redisKVCache) Destory() {
	if c.destroy != nil {
		c.destroy()
	}
}

type redisCacheFactory struct {
	redisPool redis_pkg.RedisPool
}

func NewRedisCacheFactory(redisPool redis_pkg.RedisPool) CacheFactory {
	return &redisCacheFactory{
		redisPool: redisPool,
	}
}

func (f *redisCacheFactory) NewKVCache() KVCache {
	client := f.redisPool.Get()
	destory := func() {
		f.redisPool.Put(client)
	}
	return newRedisKVCache(client, destory)
}

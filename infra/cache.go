package infra

import (
    "context"
    "time"

    "github.com/go-redis/cache/v9"
)

var RedisCache *cache.Cache

func InitCache() {
    RedisCache = cache.New(&cache.Options{
        Redis: RedisClient, // `redis.go` の RedisClient を使用
    })
}

// SetCache は Redis にデータをキャッシュする
func SetCache(key string, value interface{}, expiration time.Duration) error {
    ctx := context.Background()
    return RedisCache.Set(&cache.Item{
        Ctx:   ctx,
        Key:   key,
        Value: value,
        TTL:   expiration,
    })
}

// GetCache はキャッシュからデータを取得する
func GetCache(key string, dest interface{}) error {
    ctx := context.Background()
    err := RedisCache.Get(ctx, key, dest)
    if err == cache.ErrCacheMiss {
        return err // キャッシュがない場合
    }
    return err
}

// DeleteCache はキャッシュを削除する
func DeleteCache(key string) error {
    ctx := context.Background()
    return RedisCache.Delete(ctx, key)
}
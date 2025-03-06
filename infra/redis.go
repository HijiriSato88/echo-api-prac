package infra

import (
    "context"
    "fmt"
    "os"

    "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
    redisHost := os.Getenv("REDIS_HOST")
    redisPort := os.Getenv("REDIS_PORT")

    addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
    RedisClient = redis.NewClient(&redis.Options{
        Addr: addr,
        DB:   0, // デフォルトのDBを使用
    })

    ctx := context.Background()
    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
    }
    fmt.Println("Connected to Redis")
}

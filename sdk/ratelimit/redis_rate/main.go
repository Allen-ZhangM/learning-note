package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"time"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)
	for {
		res, err := limiter.Allow(ctx, "project:123", redis_rate.PerMinute(60))
		if err != nil {
			panic(err)
		}
		fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
		time.Sleep(500 * time.Millisecond)
	}
}

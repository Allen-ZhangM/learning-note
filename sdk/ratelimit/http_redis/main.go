package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"log"
	"net/http"
)

var limiter *redis_rate.Limiter
var ctx = context.Background()

func main() {
	initLimiter()

	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	if err := http.ListenAndServe(":8888", limitMiddleware(mux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter == nil {
			log.Println("limiter is nil")
		}
		limit, _ := limiter.Allow(ctx, r.RemoteAddr, redis_rate.PerSecond(5))
		if limit.Allowed != 1 {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	// Some very expensive database call
	w.Write([]byte("test limit"))
}

func initLimiter() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	err := rdb.FlushDB(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}

	limiter = redis_rate.NewLimiter(rdb)
}

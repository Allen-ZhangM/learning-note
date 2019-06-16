package chanratelimit

import (
	"log"
	"testing"
	"time"
)

func TestChanRateLimiter_Limit(t *testing.T) {
	rl := New(5, time.Second)
	for i := 0; i < 100; i++ {
		log.Printf("limit result: %v\n", rl.Limit())
		time.Sleep(10 * time.Millisecond)
	}
}

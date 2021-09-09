package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dgraph-io/ristretto"
)

func TestLimit(t *testing.T) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		// num of keys to track frequency, usually 10*MaxCost
		NumCounters: 100,
		// cache size(max num of items)
		MaxCost: 10,
		// number of keys per Get buffer
		BufferItems: 64,
		// !important: always set true if not limiting memory
		IgnoreInternalCost: true,
	})
	if err != nil {
		panic(err)
	}

	// put 20(>10) items to cache
	for i := 0; i < 20; i++ {
		cache.Set(i, i, 1)
	}

	// wait for value to pass through buffers
	cache.Wait()

	cntCacheMiss := 0
	for i := 0; i < 20; i++ {
		if _, ok := cache.Get(i); !ok {
			cntCacheMiss++
		}
	}
	fmt.Printf("%d of 20 items missed\n", cntCacheMiss)
}

func TestTTL(t *testing.T) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters:        100,
		MaxCost:            10,
		BufferItems:        64,
		IgnoreInternalCost: true,
	})
	if err != nil {
		panic(err)
	}

	// set item with 1s ttl
	cache.SetWithTTL("foo", "bar", 1, time.Second)

	// wait for value to pass through buffers
	cache.Wait()

	if val, ok := cache.Get("foo"); !ok {
		log.Printf("cache missing")
	} else {
		log.Printf("got foo: %v", val)
	}

	// sleep longer and try again
	time.Sleep(2 * time.Second)
	if val, ok := cache.Get("foo"); !ok {
		log.Printf("cache missing")
	} else {
		log.Printf("got foo: %v", val)
	}
}

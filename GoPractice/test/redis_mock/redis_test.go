package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-redis/redismock"
	"testing"
	"time"
)

var ctx = context.TODO()

func NewsInfoForCache(redisDB *redis.Client, newsID int) (info string, err error) {
	cacheKey := fmt.Sprintf("news_redis_cache_%d", newsID)
	info, err = redisDB.Get(ctx, cacheKey).Result()

	// info, err = call api()
	info = "test"
	err = redisDB.Set(ctx, cacheKey, info, 30*time.Minute).Err()

	s, err := redisDB.SMembers(ctx, "smber").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

	err = redisDB.Expire(ctx, "expire", time.Minute).Err()

	err = redisDB.Set(ctx, cacheKey, info, 30*time.Minute).Err()

	err = redisDB.SAdd(ctx, "sadd").Err()
	if err != nil {
		fmt.Println(err)
	}

	return
}

func TestNewsInfoForCache(t *testing.T) {
	db, mock := redismock.NewClientMock()

	newsID := 123456789
	key := fmt.Sprintf("news_redis_cache_%d", newsID)

	// mock ignoring `call api()`

	mock.ExpectGet(key).SetVal("value")
	mock.Regexp().ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))
	mock.Regexp().ExpectSMembers("smber").SetVal([]string{"1", "2"})
	mock.Regexp().ExpectExpire("expire", time.Minute).SetVal(true)
	mock.CustomMatch(func(expected, actual []interface{}) error {
		// Custom matching 对比两个数组的内容即可
		//[]interface{}{"set","key","value","ex",100}
		fmt.Println(expected)
		//[]interface{}{"set","news_redis_cache_123456789","test","ex",1800}
		fmt.Println(actual)
		return nil
	}).ExpectSet("key", "value", 100*time.Second).SetVal("set value")
	mock.Regexp().ExpectSAdd("sadd").SetErr(errors.New("FAIL"))

	_, err := NewsInfoForCache(db, newsID)
	if err == nil || err.Error() != "FAIL" {
		t.Error("wrong error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

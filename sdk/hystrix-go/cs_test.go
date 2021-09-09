package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"sync/atomic"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	hystrix.ConfigureCommand("aaa", hystrix.CommandConfig{
		Timeout:               5000,
		MaxConcurrentRequests: 20,
	})

	hystrix.ConfigureCommand("bbb", hystrix.CommandConfig{
		Timeout:               5000,
		MaxConcurrentRequests: 40,
	})

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("change conf")
		hystrix.ConfigureCommand("aaa", hystrix.CommandConfig{
			Timeout:               5000,
			MaxConcurrentRequests: 2,
		})

		hystrix.ConfigureCommand("bbb", hystrix.CommandConfig{
			Timeout:               5000,
			MaxConcurrentRequests: 4,
		})
	}()

	http.HandleFunc("/aaa", AHandle)

	http.HandleFunc("/bbb", BHandle)

	http.HandleFunc("/hello", ordinaryHandler)

	http.ListenAndServe(":8001", nil)
}

func TestClient(t *testing.T) {
	result := make(chan string)

	for i := 0; i < 50; i++ {
		go func() {
			run(result, "aaa")
			run(result, "bbb")
		}()
	}

	for {
		select {
		case r := <-result:
			fmt.Print(r)
		default:
		}
	}
}

func TestHystrix(t *testing.T) {
	hystrix.ConfigureCommand("t", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  20,
		RequestVolumeThreshold: 20,
		SleepWindow:            5000,
		ErrorPercentThreshold:  1,
	})

	var count_i, limet_i, circuit_i int32

	go func() {
		for range time.NewTicker(time.Second).C {
			fmt.Printf("%v  send count = %d, limet_i=%d, circuit_i=%d\n", time.Now().String(), count_i, limet_i, circuit_i)
			atomic.SwapInt32(&count_i, 0)
			atomic.SwapInt32(&limet_i, 0)
			atomic.SwapInt32(&circuit_i, 0)
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			//producer := time.NewTicker(1*time.Millisecond)
			//for range producer.C {
			for {
				hystrix.Go("t", func() error {
					atomic.SwapInt32(&count_i, count_i+1)
					return nil
				}, func(err error) error {
					switch err {
					case hystrix.ErrMaxConcurrency:
						atomic.SwapInt32(&limet_i, limet_i+1)
					case hystrix.ErrCircuitOpen:
						atomic.SwapInt32(&circuit_i, circuit_i+1)
					default:
						fmt.Println("err:=", err)
					}
					return nil
				})
			}
		}()
	}

	select {}

}

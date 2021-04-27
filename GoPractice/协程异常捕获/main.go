package main

import (
	"awesomeProject/协程异常捕获/xsync"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ids := []int64{1, 2}
	total, err := MCount(ctx, ids)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)
}

/**
recover无法捕获协程中的panic
*/
func cachePanic() {
	// 希望捕获所有所有 panic
	defer func() {
		r := recover()
		fmt.Println("recover:", r)
	}()

	// 启动新协程
	go func() {
		panic(123)
	}()
	// 等待一下，不然协程可能来不及执行
	time.Sleep(1 * time.Second)
	fmt.Println("这条消息打印不出来")
}

func cachePanicGroup() {
	// 希望捕获所有所有 panic
	defer func() {
		r := recover()
		fmt.Println("recover:", r)
	}()
	g := xsync.NewPanicGroup()
	f1 := func() {
		panic(123)
	}
	f2 := func() {
		panic(1234)
	}
	g.Go(f1)
	g.Go(f2)
	err := g.Wait(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	// 也可以链式调用
	//err = xsync.NewPanicGroup().Go(f1).Go(f2).Wait(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}
	// 等待一下，不然协程可能来不及执行
	time.Sleep(1 * time.Second)
	fmt.Println("这条消息打印不出来")
}

// MCount /**
/**
一个需要主意的设计就是Go方法的入参。有人喜欢设计成func(context.Context) error，我想了一下，觉得这种设计不好。

f是由PanicGroup执行的，如果需要传入一个ctx，那势必需要在PanicGroup中保存当前ctx，就是不被推荐的。当然，有人会说你可以在调Go方法的时候先把f存下来，等到调Wait方法的时候再统一起协程。这样做确实可以不用保存ctx了，却要保存所有的f，我也是不推荐这种做法。

f返回了一个error，这看起来是通用设计，返回报错总是天经地义的。可是，f是由PanicGroup 执行的，如果要返回错误，PanicGroup就需要跟踪哪些协程有报错、哪些没有报错，不同的报错如何跟协程对应。这会让PanicGroup的设计变得非常复杂。

基于以上两点，我建议将f设计成func()，完全没有参数，PanicGroup只跟踪panic就好了，简单明了。那如何给协程传参并处理返回结果呢？使用闭包！

这是一个实际的传参示例：
*/
func MCount(ctx context.Context, ids []int64) (total map[int]int64, err error) {
	m := sync.Mutex{}
	g := xsync.NewPanicGroup()
	total = make(map[int]int64)

	chunkSize := 50
	for i := 0; i < len(ids); i += chunkSize {
		begin := i
		end := i + chunkSize
		if end > len(ids) {
			end = len(ids)
		}

		g.Go(func() {
			ts, err := mCount(ctx, ids[begin:end])
			if err != nil {
				// log error
				return
			}

			m.Lock()
			defer m.Unlock()
			for k, v := range ts {
				total[k] = v
			}
		})
	}

	err = g.Wait(ctx)
	return
}

func mCount(ctx context.Context, slice []int64) ([]int64, error) {
	return slice, nil
}

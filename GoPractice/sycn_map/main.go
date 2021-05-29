package main

import (
	"fmt"
	"sync"
	"time"
)

//参考：https://juejin.cn/post/6844903895227957262
func main() {
	//writeMap()
	writeSyncMap()
}

func writeMap() {
	m := map[int]int{1: 1}
	//写同一个map就不行，哪怕是不同的key
	go do(m)
	go do2(m)

	time.Sleep(1 * time.Second)
	fmt.Println(m)
}

func do(m map[int]int) {
	i := 0
	for i < 10000 {
		m[i] = 1
		i++
	}
}

func do2(m map[int]int) {
	i := 0
	for i < 10000 {
		m[100000+i] = 1
		i++
	}
}

func writeSyncMap() {
	// 文章里面有问题，在这里应该用指针类型
	m := &sync.Map{}
	m.Store(1, 1)
	go doSync(m)
	go doSync(m)

	time.Sleep(1 * time.Second)
	fmt.Println(m.Load(1))
}

func doSync(m *sync.Map) {
	i := 0
	for i < 10000 {
		m.Store(1, 1)
		i++
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//普通chan，无顺序输出，先取得的值先输出
	channel()
	//嵌套chan，可以按照添加顺序输出
	//channelN()

}

func channelN() {

	resultCh := make(chan chan string, 4)

	go replayN(resultCh)

	wg := sync.WaitGroup{}

	startTime := time.Now()

	operationN2(resultCh, "aaa", &wg)
	operationN2(resultCh, "bbb", &wg)
	operationN1(resultCh, "ccc", &wg)
	operationN2(resultCh, "ddd", &wg)
	operationN2(resultCh, "eee", &wg)
	wg.Wait()
	endTime := time.Now()
	fmt.Printf("Process time %s", endTime.Sub(startTime))

}

func operationN1(resultCh chan chan string, str string, wg *sync.WaitGroup) {
	wg.Add(1)

	c := make(chan string)

	resultCh <- c

	go func(str string) {
		time.Sleep(time.Second * 1)
		c <- "operation1:" + str
		wg.Done()
	}(str)

}

func operationN2(resultCh chan chan string, str string, wg *sync.WaitGroup) {
	wg.Add(1)

	c := make(chan string)

	resultCh <- c

	go func(str string) {
		time.Sleep(time.Second * 3)
		c <- "operation2:" + str
		wg.Done()
	}(str)

}

func replayN(resultCh chan chan string) {
	for {
		r := <-resultCh
		fmt.Println(<-r)
	}
}

func channel() {
	resultCh := make(chan string)
	//开一个gotoutine 接受所有返回值并打印
	go replay(resultCh)
	//使用waitgroup 等待一下所有gorountie执行完毕，记录时间
	wg := sync.WaitGroup{}

	startTime := time.Now()

	//operation1内部sleep 1秒
	//operation2内部sleep 2秒
	//如果是同步执行下列调用需要 8秒左右
	//目前用异步调用 理论上只需要2秒
	//但于丹的问题是 不能实现先进先出的需求
	operation2(resultCh, "aaa", &wg)
	operation2(resultCh, "bbb", &wg)
	operation1(resultCh, "ccc", &wg)
	operation1(resultCh, "ddd", &wg)
	operation2(resultCh, "eee", &wg)
	wg.Wait()
	endTime := time.Now()
	fmt.Printf("Process time %s", endTime.Sub(startTime))
}

func replay(resultCh chan string) {
	for {
		fmt.Println(<-resultCh)
	}
}

func operation1(resultCh chan string, str string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(str string, wg *sync.WaitGroup) {
		time.Sleep(time.Second * 1)
		resultCh <- "operation1:" + str
		wg.Done()
	}(str, wg)
}

func operation2(resultCh chan string, str string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(str string, wg *sync.WaitGroup) {
		time.Sleep(time.Second * 2)
		resultCh <- "operation2:" + str
		wg.Done()
	}(str, wg)
}

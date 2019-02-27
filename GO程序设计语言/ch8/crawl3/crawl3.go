package main

import (
	"../../ch5/links"
	"fmt"
	"os"
)

func main() {
	workList := make(chan []string)  //可能重复的url列表
	unseenLinks := make(chan string) //去重复的url列表

	//向任务列表中添加命令行参数
	go func() { workList <- os.Args[1:] }()

	//创建20个goroutine获取每个不可见连接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}

	//主goroutine对url列表去重
	//把没有爬取过的条目发送给爬虫程序
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

// 爬取 url 页面的所有链接并返回
func crawl(url string) (subURLs []string) {
	fmt.Println(url)
	subURLs, err := links.Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return
}

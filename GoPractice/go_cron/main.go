package main

import (
	"github.com/robfig/cron"
	"log"
)

const (
	every5sec  = "*/5 * * * * ?"
	every5sec_ = "@every 5s"
	every1min_ = "0 */1 * * * ?"
)

func main() {
	log.Println("Starting...")

	c := cron.New(cron.WithSeconds())
	c.AddFunc(every5sec, func() {
		log.Println("Run 每隔5秒执行一次...")
	})
	c.AddFunc(every5sec_, func() {
		log.Println("Run 从启动时算起 每隔5秒执行一次...")
	})
	c.AddFunc(every1min_, func() {
		log.Println("Run 每隔1分钟执行一次...")
	})

	c.Start()
	defer c.Stop()

	select {}
}

package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"learning-note/nsq-go/agent"
	"learning-note/nsq-go/consumer"
)

func initLogger() {
	logs.SetLevel(7)
	// default 4 for calling with beego.Debug()/Warn()...
	// while 3 is correct for calling with logs.Debug()/Warn()...
	logs.SetLogFuncCallDepth(3)

	filePath := beego.AppConfig.String("log::logfile")
	logConfig := fmt.Sprintf(`{"filename":"%s", "perm":"0644"}`, filePath)
	logs.SetLogger("file", logConfig)
	beego.BeeLogger.DelLogger("console")
}

func main() {
	//初始化生产者
	pc := &agent.ProducerConf{
		Addr:         "192.168.194.128:4150",
		DoneChanSize: 100000,
		HandleAsyncErrFunc: func(err error, topic string, msg []byte) {
			fmt.Println(err, topic, string(msg))
			logs.Error("failed to send msg , err:  %s , topic: %s , msg : %s", err, topic, msg)
		},
	}
	agent.InitNsqd(pc)

	//初始化消费者
	c := &consumer.Conf{
		Topic:    "test",
		Channel:  "c1",
		AddrNsqd: "192.168.194.128:4150",
	}
	consumer.InitNsqConsumer(*c)
}

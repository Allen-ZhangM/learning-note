package main

import (
	"flag"
	"fmt"
	"learning-note/sdk/kafka/producer"
	"strings"
	"sync"
)

var topic = flag.String("topic", "", "发送的topic")
var value = flag.String("value", "", "发送的内容byte")
var brokers = flag.String("brokers", "", "brokers逗号分隔")
var count = flag.Int("count", 1, "发送条数")

var wg sync.WaitGroup

//启动命令参考： ./cmd_producer -topic topic  -brokers 192.168.194.128:9092
func main() {
	flag.Parse()

	fmt.Printf("params: topic:%s, brokers:%s, count:%d, value:%s \n", *topic, *brokers, *count, *value)

	p, err := producer.InitProducer(&producer.ProducerConf{
		Addrs: strings.Split(*brokers, ","),
		SuccessFunc: func(info *producer.CallbackInfo) {
			fmt.Printf("send success: ;topic:%s;msg:%s;offset:%d;partition:%d\n", info.ProducerMessage.Topic, info.ProducerMessage.Value, info.ProducerMessage.Offset, info.ProducerMessage.Partition)
			wg.Done()
		},
		ErrorFunc: func(info *producer.CallbackInfo) {
			fmt.Printf("send err:%v ;topic:%s;msg:%s;offset:%d;partition:%d\n", info.Err, info.ProducerMessage.Topic, info.ProducerMessage.Value, info.ProducerMessage.Offset, info.ProducerMessage.Partition)
			wg.Done()
		},
	})
	if err != nil {
		panic(err)
	}
	wg.Add(*count)
	i := 0
	for i = 0; i < *count; i++ {
		//单条耗时：183.916µs
		p.PushMsg(&producer.ProducerMessage{
			Topic: *topic,
			Value: []byte(*value),
		})
	}

	wg.Wait()
	fmt.Println("execute success.num:", i)

}

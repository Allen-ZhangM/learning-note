package main

import (
	"flag"
	"fmt"
	"learning-note/sdk/kafka/producer"
	"strings"
	"sync"
	"sync/atomic"
)

var topic = flag.String("topic", "", "发送的topic")
var value = flag.String("value", "", "发送的内容byte")
var brokers = flag.String("brokers", "", "brokers逗号分隔")
var count = flag.Int("count", 1, "发送条数")
var producerCount = flag.Int("producer_count", 1, "生产者数")

var wg sync.WaitGroup

//启动命令参考： ./cmd_producer -topic topic  -brokers 192.168.194.128:9092
func main() {
	flag.Parse()

	fmt.Printf("params: topic:%s, brokers:%s, count:%d, value:%s ,producer_count:%d \n", *topic, *brokers, *count, *value, *producerCount)

	producers := make([]*producer.KafkaProducer, 0, *producerCount)
	for i := 0; i < *producerCount; i++ {
		p, err := producer.InitProducer(&producer.ProducerConf{
			Addrs: strings.Split(*brokers, ","),
			SuccessFunc: func(info *producer.CallbackInfo) {
				//fmt.Printf("send success: ;topic:%s;msg:%s;offset:%d;partition:%d\n", info.ProducerMessage.Topic, info.ProducerMessage.Value, info.ProducerMessage.Offset, info.ProducerMessage.Partition)
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
		producers = append(producers, p)
	}

	fmt.Printf("producers_num:%d \n", len(producers))

	wg.Add(*count * *producerCount)
	var n int32 = 0
	for index := range producers {
		go func(index int) {
			for i := 0; i < *count; i++ {
				//单条耗时：183.916µs
				producers[index].PushMsg(&producer.ProducerMessage{
					Topic: *topic,
					Value: []byte(*value),
				})
				atomic.AddInt32(&n, 1)
			}
		}(index)
	}

	wg.Wait()
	fmt.Println("execute success.num:", n)

}

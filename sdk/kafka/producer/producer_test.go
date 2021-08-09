package producer

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {

	//peer_session   biz_report   lsm_report  vv_session
	topic := "test"
	topic2 := "test2"
	topic3 := "test3"
	topic4 := "test4"
	addrs := []string{"192.168.194.128:9092", "192.168.194.128:9093", "192.168.194.128:9094"}
	p, err := InitProducer(&ProducerConf{
		Addrs: addrs,
		SuccessFunc: func(info *CallbackInfo) {
			//fmt.Println("send success ", info.producerMessage.Topic)
			fmt.Printf("send success: ;topic:%s;msg:%s;offset:%d", info.ProducerMessage.Topic, info.ProducerMessage.Value, info.ProducerMessage.Offset)
		},
		ErrorFunc: func(info *CallbackInfo) {
			fmt.Printf("send err:%v ;topic:%s;msg:%s;offset:%d", info.Err, info.ProducerMessage.Topic, info.ProducerMessage.Value, info.ProducerMessage.Offset)
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	i := 0

	go func() {
		for range time.NewTicker(time.Second).C {
			fmt.Printf("%v  send count = %d\n", time.Now().String(), i)
		}
	}()

	for {
		//单条耗时：183.916µs
		p.PushMsg(&ProducerMessage{
			Topic: topic,
			Value: []byte("test" + time.Now().String() + " 序号:" + strconv.Itoa(i)),
		})
		p.PushMsg(&ProducerMessage{
			Topic: topic2,
			Value: []byte("test2" + time.Now().String() + " 序号:" + strconv.Itoa(i)),
		})
		p.PushMsg(&ProducerMessage{
			Topic: topic3,
			Value: []byte("test3" + time.Now().String() + " 序号:" + strconv.Itoa(i)),
		})
		p.PushMsg(&ProducerMessage{
			Topic: topic4,
			Value: []byte("test4" + time.Now().String() + " 序号:" + strconv.Itoa(i)),
		})
		i++
		time.Sleep(time.Millisecond)
	}

}

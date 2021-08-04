package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pt "learning-note/GoPractice/rpc/grpc-proto/proto"
	"time"
)

const (
	post_c = "127.0.0.1:18881"
)

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second)
	defer cancelFunc()

	//客户端连接服务器
	conn, err := grpc.DialContext(ctx, post_c, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
		return
	}
	defer conn.Close()

	//获得grpc句柄
	c := pt.NewHelloServerClient(conn)
	//远程调用 SayHello接口
	r1, err := c.SayHello(ctx, &pt.HelloRequest{Name: "panda"})
	if err != nil {
		fmt.Println("cloud not get Hello server ..", err)
		return
	}
	fmt.Println("HelloServer resp: ", r1.Message)

	//远程调用 GetHelloMsg接口
	r2, err := c.GetHelloMsg(ctx, &pt.HelloRequest{Name: "panda"})
	if err != nil {
		fmt.Println("cloud not get hello msg ..", err)
		return
	}
	fmt.Println("HelloServer resp: ", r2.Msg)

	select {
	case <-ctx.Done():
		fmt.Println("call successfully!!!")
		return
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("timeout!!!")
		return
	}

}

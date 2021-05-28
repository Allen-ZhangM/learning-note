package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pt "learning-note/rpc/grpc-proto/proto"
	"testing"
)

/**
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkInvokeCall
BenchmarkInvokeCall-4   	    4083	    245748 ns/op
*/
func BenchmarkInvokeCall(b *testing.B) {
	ctx := context.Background()

	//客户端连接服务器
	conn, err := grpc.DialContext(ctx, post_c, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
		return
	}
	defer conn.Close()

	//获得grpc句柄
	c := pt.NewHelloServerClient(conn)

	for i := 0; i < b.N; i++ {
		//远程调用 SayHello接口
		_, err := c.SayHello(ctx, &pt.HelloRequest{Name: "panda"})
		if err != nil {
			fmt.Println("cloud not get Hello server ..", err)
			return
		}
	}

}

/**
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkInvokeClient
BenchmarkInvokeClient-4   	    3840	    261126 ns/op
*/
func BenchmarkInvokeClient(b *testing.B) {
	ctx := context.Background()

	//客户端连接服务器
	conn, err := grpc.DialContext(ctx, post_c, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
		return
	}
	defer conn.Close()
	for i := 0; i < b.N; i++ {
		//获得grpc句柄
		c := pt.NewHelloServerClient(conn)

		//远程调用 SayHello接口
		_, err := c.SayHello(ctx, &pt.HelloRequest{Name: "panda"})
		if err != nil {
			fmt.Println("cloud not get Hello server ..", err)
			return
		}
	}

}

/**
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkInvokeConn
BenchmarkInvokeConn-4   	    1226	    954105 ns/op
*/
func BenchmarkInvokeConn(b *testing.B) {
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		//客户端连接服务器
		conn, err := grpc.DialContext(ctx, post_c, grpc.WithInsecure())
		if err != nil {
			fmt.Println("连接服务器失败", err)
			return
		}

		//获得grpc句柄
		c := pt.NewHelloServerClient(conn)

		//远程调用 SayHello接口
		_, err = c.SayHello(ctx, &pt.HelloRequest{Name: "panda"})
		if err != nil {
			fmt.Println("cloud not get Hello server ..", err)
			return
		}
		conn.Close()
	}

}

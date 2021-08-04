package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pt "learning-note/GoPractice/rpc/grpc-proto/proto"
	"testing"
	"time"
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

func TestContext(t *testing.T) {
	num := 10000
	ctx, cancelF := context.WithTimeout(context.Background(), time.Minute)
	defer cancelF()
	ctx2, cancelF2 := context.WithTimeout(context.Background(), time.Minute)
	defer cancelF2()
	ctx3, cancelF3 := context.WithTimeout(context.Background(), time.Minute)
	defer cancelF3()
	//客户端连接服务器
	conn, err := grpc.DialContext(ctx, post_c, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
		return
	}
	defer conn.Close()

	//获得grpc句柄
	c := pt.NewHelloServerClient(conn)

	for i := 0; i < num; i++ {
		//远程调用 SayHello接口
		_, err := c.SayHello(ctx, &pt.HelloRequest{Name: "panda"})
		if err != nil {
			fmt.Println("cloud not get Hello server ..", err)
			return
		}
	}

	//客户端连接服务器
	conn, err = grpc.DialContext(ctx2, post_c, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
		return
	}
	defer conn.Close()
	for i := 0; i < num; i++ {
		//获得grpc句柄
		c := pt.NewHelloServerClient(conn)

		//远程调用 SayHello接口
		_, err := c.SayHello(ctx2, &pt.HelloRequest{Name: "panda"})
		if err != nil {
			fmt.Println("cloud not get Hello server ..", err)
			return
		}
	}

	for i := 0; i < num; i++ {
		//客户端连接服务器
		conn, err := grpc.DialContext(ctx3, post_c, grpc.WithInsecure())
		if err != nil {
			fmt.Println("连接服务器失败", err)
			return
		}

		//获得grpc句柄
		c := pt.NewHelloServerClient(conn)

		//远程调用 SayHello接口
		_, err = c.SayHello(ctx3, &pt.HelloRequest{Name: "panda"})
		if err != nil {
			fmt.Println("cloud not get Hello server ..", err)
			return
		}
		conn.Close()
	}
	d1, _ := ctx.Deadline()
	d2, _ := ctx2.Deadline()
	d3, _ := ctx2.Deadline()
	fmt.Println("ctx1:", d1)
	fmt.Println("ctx2:", d2)
	fmt.Println("ctx3:", d3)

}

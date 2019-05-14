package main

import (
	"../../jsonrpc"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	c, e := net.Dial("tcp", ":1234")
	if e != nil {
		panic(e)
	}
	client := jsonrpc.NewClient(c)

	var result *float64
	e = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if e != nil {
		log.Println(e)
		panic(e)
	}
	fmt.Println(*result)

}

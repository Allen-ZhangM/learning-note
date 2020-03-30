package main

import (
	demo "../../jsonrpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {

	conn, _ := net.Dial("tcp", ":1234")

	client := jsonrpc.NewClient(conn)

	var result float64
	err := client.Call("CalculateService.Div", demo.Args{3, 5}, &result)
	fmt.Println(result, err)

}

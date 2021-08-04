package main

import (
	"fmt"
	"learning-note/GoPractice/rpc/jsonrpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {

	conn, _ := net.Dial("tcp", ":1234")

	client := jsonrpc.NewClient(conn)

	var result float64
	err := client.Call("CalculateService.Div", demorpc.Args{3, 5}, &result)
	fmt.Println(result, err)

}

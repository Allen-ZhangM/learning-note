package main

import (
	"../../jsonrpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		panic(e)
	}

	for {
		c, e := l.Accept()
		if e != nil {
			log.Println("Accept err")
			continue
		}
		go jsonrpc.ServeConn(c)
	}

}

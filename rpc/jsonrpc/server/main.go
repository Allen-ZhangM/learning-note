package main

import (
	"../../jsonrpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(demorpc.CalculateService{})
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		panic(e)
	}

	for {
		c, e := l.Accept()
		if e != nil {
			log.Println("accept err :", e)
			continue
		}
		go jsonrpc.ServeConn(c)
	}
}

//开启服务
//linux 系统的telnet
// telnet localhost 1234
//{"method":"CalculateService.Div","params":[{"A":1,"B":2}],"id":1}

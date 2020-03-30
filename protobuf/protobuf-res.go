package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello:" + request.GetValue()
	return nil
} //rpc 调用要求，第二个参数必须是指针，并且返回值 必须是 error

func main() {
	// 整个步骤是1. listen 2. 建立连接， 3.rpc 调用连接
	//
	rpc.RegisterName("HelloService", new(HelloService))

	listenner, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp error:", err)
	}

	conn, err := listenner.Accept()
	if err != nil {
		log.Fatal("Accept error", err)
	}
	rpc.ServeConn(conn)
}

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
 *  这里 nc -l 1234 启动服务端
 *  客户端启动后，服务端 会 收到 json 数据
 *  echo -e '{"method":"HelloService.Hello", "params": ["hello"], "id":1]}' | nc localhost 1234
 */
func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.dialog:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

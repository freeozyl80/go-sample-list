package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234") //拨打
	if err != nil {
		log.Fatal("dialog:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

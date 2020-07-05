package main

import (
	"fmt"
	"time"
)
func async(data string) {

 	d := time.Duration(time.Second*2)
	ticker := time.NewTicker(d)
	tChan := make(chan interface{}, 1)
	select {
	case <- ticker.C:
		fmt.Println("超时")
	case tChan <- data:
		fmt.Println("已赋值")
	}
}

func main() {
	defer func() {
		fmt.Println("Test pending for queue")
	}()

    time.Sleep(3 * time.Second)
	go async("start")
}
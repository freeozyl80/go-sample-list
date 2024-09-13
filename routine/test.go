package main

import (
	"fmt"
	"time"
)

func main() {

	donec := make(chan bool, 1)
	// close(donec)
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("timer")
		case <-donec:
		}
	}
}

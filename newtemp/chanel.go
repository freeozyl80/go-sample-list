// range 才能显示，这是个坑
package main

import (
	"fmt"
	"time"
)
func main() {
	go func() {
		time.Sleep(1 * time.Second)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}

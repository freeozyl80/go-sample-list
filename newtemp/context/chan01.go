package main

import (
	"fmt"
	"time"
	"runtime"
	
)

func main() {
 runtime.GOMAXPROCS(2)
 ch := make(chan int) // 无缓冲
 // 队列的概念，我日
 go goRoutineA(ch)
 go goRoutineB(ch)
 ch <- 1
 ch <- 2

 time.Sleep(time.Second * 2)
}

func goRoutineA(ch chan int) {
 v := <-ch
 fmt.Printf("A received data: %d\n", v)
}

func goRoutineB(ch chan int) {
 v := <-ch
 fmt.Printf("B received data: %d\n", v)
}

package main

import (
	"fmt"
	"time"
)
func main() {
 ch := make(chan int, 3) // 有缓冲

 // 都不会阻塞
 ch <- 1
 ch <- 2
 ch <- 3
    // 会阻塞，被挂起到 sendq 中
 go func() {
  ch <- 4
 }()

 // 只是为了debug
 var a int
 fmt.Println(a)

 go goRoutineA(ch)
 go goRoutineA(ch)
 go goRoutineB(ch)
 go goRoutineB(ch) // 猜猜这里会被挂起吗？

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

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	userCount := 10
	ch := make(chan int, 50)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("xxx")
			for d := range ch { // range  可以 使得 一些 没有赋值 的 ch 被 过滤掉
				fmt.Printf("go func: %d, time: %d\n", d, time.Now().Unix())
				time.Sleep(time.Second * time.Duration(d))
			}
		}()
	}

	for i := 0; i < 2; i++ {
		ch <- 1
		ch <- 3
		//time.Sleep(time.Second)
	}

	close(ch)
	wg.Wait()
}

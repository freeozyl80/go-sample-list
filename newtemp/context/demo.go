package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

var wg sync.WaitGroup

func run(task string) {
	fmt.Println(task, "start ...")
	time.Sleep(time.Second * 2)

	wg.Done()
}

func main() {
	wg.Add(3)
	for i := 1; i <4; i++ {
		taskName := "task" + strconv.Itoa(i)
		go run(taskName)
	}

	wg.Wait()
	fmt.Println("所有任务结束")
}
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32          //计数器
	wg      sync.WaitGroup //信号量
)

func main() {
	fmt.Println(counter)
	threadNum := 100
	//1. 五个信号量
	wg.Add(threadNum)
	//2.开启5个线程
	for i := 0; i < threadNum; i++ {
		go incCounter(i)
	}
	//3.等待子线程结束
	wg.Wait()
	fmt.Println(counter)
}

func incCounter(index int) {
	defer wg.Done()
	spinNum := 0
	for {
		//2.1原子操作
		fmt.Printf("counter in thread is %d", counter)
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	fmt.Printf("thread,%d,spinnum,%d\n", index, spinNum)
}

/*
读者可以把线程数改为10000或者更多会发现输出thread,5329,spinnum,1其中1说明该线程尝试了两个CAS操作，第二次才成功。
*/

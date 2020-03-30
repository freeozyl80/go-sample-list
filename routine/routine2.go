package main

import (
	"sync"
)

func main() {
	var mu sync.Mutex

	mu.Lock()

	go func() {
		println("hi")
		mu.Unlock()
	}()

	mu.Lock()
}

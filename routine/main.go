package main

import (
	"fmt"
	"reflect"
	"routine/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(10000000000*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	fmt.Println(reflect.TypeOf(all))

	p.Publish("Hello World")
	p.Publish("Hello golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(30 * time.Second)
}

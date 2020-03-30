package main

import (
	"fmt"
)

type f interface {
	Error(args ...interface{})
}
type TB struct {
	f
}

func (tb *TB) Error(args ...interface{}) {
	fmt.Println("TB.Error disabled!")
}
func main() {
	fmt.Println("Hello World")
	var tb f = new(TB)
	fmt.Printf("%+v\n", tb)
	tb.Error("Hello")
}

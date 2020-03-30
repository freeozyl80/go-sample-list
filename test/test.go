package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"test/try"
)

func main() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(bytes.Equal(b, make([]byte, c)))

	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))

	try.PrintSecond()
}

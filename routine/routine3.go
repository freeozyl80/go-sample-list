package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "HI"
	done <- true
}
func main() {
	go aGoroutine()
	<-done
	println(msg)
}

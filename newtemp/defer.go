package main

import "fmt"

func f1() (result int) {
    defer func() {
        result++
    }()
    return 0
}
func f2() (r int) {
     t := 5
     defer func() {
       t = t + 5
     }()
     return t
}
func f3() (r int) {
    defer func(r int) {
          r = r + 5
    }(r)
    return 1
}
func main() {

	var test = f1()
	fmt.Println(test)
}

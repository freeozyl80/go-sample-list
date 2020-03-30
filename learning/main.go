package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = arr1

	arr1[0] = 0

	fmt.Println(arr1)
	fmt.Println(arr2)

	var p = &arr1
	for i, v := range p {
		fmt.Println(i, v)
	}

	var slice1 = []int{1, 2, 3}
	var slice2 = slice1

	slice1[0] = 2
	fmt.Println(slice2)

	availableParamentF(arr1, slice1[0])
}

func availableParamentF(p ...interface{}) {
	fmt.Println(p...)
}

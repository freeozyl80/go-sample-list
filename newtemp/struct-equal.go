// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
  A string
  B int64
  C [3]int
}

type T2 struct {
  A string
  B int64
  C []int
}

func main() {
	temp1 := T2{A:"zhang",B: 64, C: []int{1,2,3}}
	fmt.Println(temp1)
	
	
	temp2 := T2{A:"zhang",B: 64, C: []int{1,2,4}}
	fmt.Println(temp2)
	
	s := reflect.ValueOf(&temp2).Elem()
	typeOfT := s.Type()
	
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	
	fmt.Println(typeOfT)
	
	if reflect.DeepEqual(temp1, temp2) {
		fmt.Printf("测试结果： %s", "equal")
	} else {
		fmt.Printf("测试结果： %s", "unqual")
	}
}


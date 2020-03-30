package main

import (
	"fmt"
	"reflect"
)

type A struct {
	UID int `json:"uid"`
}
type OwnString struct {
	A
	UUID int `json:"uid"`
}

type T struct {
	A int
	B string
}

func main() {
	var ows1 = OwnString{}
	ows1.UUID = 1
	ows1.A.UID = 1

	//fmt.Println(ows1)

	var f interface{}
	f = ows1

	fmt.Println(reflect.TypeOf(f).Name())

	fmt.Println(f)

								                                          
	  := reflect.TypeOf(f)
	getValue := reflect.ValueOf(f)
	getElm := reflect.ValueOf(&f).Elem()

	fmt.Println(getElm)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)

		value := getValue.Field(i).Interface()
		fmt.Println("d----")
		fmt.Printf("%T", getElm)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	fmt.Println("----")
	fmt.Printf("%T", getElm)
	fmt.Println()

	var tt interface{}
	tt = T{23, "skidoo"}
	getElm.Set(reflect.ValueOf(tt))
	fmt.Println(f)

	fmt.Println("============== cut ==============")

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

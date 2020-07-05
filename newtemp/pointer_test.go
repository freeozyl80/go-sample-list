package main

import (
	"testing"
	"unsafe"
)
var ptrSize uintptr

func init() {
	ptrSize = unsafe.Sizeof(uintptr(1))
}
type SType struct {
	b [32]byte
}

func BenchmarkAligned(b *testing.B) {
	x := SType{}
	address := uintptr(unsafe.Pointer(&x.b)) + 1
	if address%ptrSize ==0 {
		b.Error("Not aligned address")
	}
	tmp := (*int64)(unsafe.Pointer(address))
	b.ResetTimer()

	for i:=0; i<b.N; i++ {
		*tmp = int64(i)
	}
}
package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {

	h := md5.New()
	io.WriteString(h, "first")
	io.WriteString(h, "second")

	test := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(test)

	data1 := []byte("firstsecond")
	// data2 := fmt.Sprintf("%x", md5.Sum(data1))
	// data := []byte(data2)
	fmt.Printf("%x\n", md5.Sum(data1))
	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
}

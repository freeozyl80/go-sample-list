## 重读 GO 语言：

* 数组是 值语义，一个数组变量即是整个数组： 当一个数组被赋值或者传递的时候，实际上会复制整个数组； 

  另外：

  ``` code
  var a = [...]int{1,2,3}
  var b = &a
  b[0]  //1
  ```

* 

```code
  slice := make([]int, 0, 100)
  hash := make(map[int]bool, 10)
  ch := make(chan int, 5)
```

* 对切片本身进行赋值和参数传递时，和数组指针的操作方式类似

* go的接口之间的 类型转换十分灵活：    xxx = x.(type)
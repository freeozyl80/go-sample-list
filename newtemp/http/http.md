# http transaction 模式

http 请求报文 和 响应 报文是 如何建立的：

* Http 协议是 transaction-driven

#### http close 模式

对于每一个请求，有且只有一个响应：

```
[CON1][REQ1]...[RESP1][CLO1]
```

####  Keep-alive 模式

Content-length

!> 由于 HTTP 协议的 transactional 属性，有了改进的方法。对于两个连续的 transactions，server 在第一次响应后不会马上关闭连接。


#### pipeling模式

!> client 不等接受第一个响应后就发送第二个请求



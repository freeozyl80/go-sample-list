队列：

  A 一级缓存，自实现队列
  B 需要链接的知名中间件
  	b1 不支持持久化
  	b2 自身支持持久化 
  	  b1.1 是否支持集群


  A:

   orederKey
   bufferSize
   HanleTime

   Interface IQueue
   	push
   	pop

   QueueTimeoutResp
   	Timeout
   	Response

   QueueRequest
   	RequId
   	Order
   	AccessTime
   	ResponseChan

   Queue
   	mapLock
   	RequestChan
   	RequestMap
   	Queue

  Q:
  	1. 试试是不是会阻塞

RabbitMQ:

TTL 消息存活时间



1.延时队列：
延时队列的元素都是带有时间属性的

DelayQueue 队列
lock 锁
priorityqueue 优先级队列
treadleader 标记是否有线程排队
condition 
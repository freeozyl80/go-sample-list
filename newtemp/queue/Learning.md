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


https://learnku.com/articles/41728
M 何时创建：没有足够的 M 来关联 P 并运行其中的可运行的 G。比如所有的 M 此时都阻塞住了，而 P 中还有很多就绪任务，就会去寻找空闲的 M，而没有空闲的，就会去创建新的 M。
package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // chan interface{} 相当于一种类型
	topicFunc  func(v interface{}) bool //知识点 type 相当与 函数的 para 和 return
)

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub) //知识点 map[chan]
	close(sub)
}

func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1) //队列
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait() //阻塞 等 0
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub) //知识点 chan 是可以被 close的
	}
}

func (p *Publisher) sendTopic(
	sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup,
) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		fmt.Println("!")
		return
	}
	select {
	case sub <- v:
		fmt.Println("有")
	case <-time.After(10 * time.Second):
		fmt.Println("无")
	}
}

//结合特性2，每次 select 都会对所有通信表达式求值，因此可通过 time.After简洁实现定时器功能，并且定时任务可通过 done channel 停止:
// 因此，case <- timer.After(time.Second) 不应该解释为每一秒执行一次，而是其它 case 如果有一秒都没有执行，那么就执行这个 case。
// waitGroup.Add  添加一个 计数器
// waitGroup.Done 减少一个 计数器
// waitGroup.Wait 等待所有计数器 清0
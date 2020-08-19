package msgpool

import (
	"sync"
)
type Message struct{
	Count int64
}

type messagePool struct {
	pool *sync.Pool
}
var msgPool = &messagePool{
    pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
}
func Instance() *messagePool {
    return msgPool
}
// 往消息池里添加消息
func (m *messagePool) AddMsg(msg *Message) {
    m.pool.Put(msg)
}
// 从消息池里获取消息
func (m *messagePool) GetMsg() *Message {
    return m.pool.Get().(*Message)
}

// func Instance() *messagePool {
//     // 在匿名函数中实现初始化逻辑，Go语言保证只会调用一次
//     once.Do(func() {
//         msgPool = &messagePool{
//             // 如果消息池里没有消息，则新建一个Count值为0的Message实例
//             pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
//         }
//     })
//     return msgPool
// }
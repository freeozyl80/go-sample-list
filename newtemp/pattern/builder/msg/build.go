package msg

import (
	"sync"
)

type builder struct {
  once *sync.Once
  msg *Message
}

func Builder() *builder {
    return &builder{
        once: &sync.Once{},
        msg: &Message{Header: &Header{}, Body: &Body{}},
    }
}
// 以下是对Message成员对构建方法
func (b *builder) WithSrcAddr(srcAddr string) *builder {
    b.msg.Header.SrcAddr = srcAddr
    return b
}
func (b *builder) WithSrcPort(srcPort uint64) *builder {
    b.msg.Header.SrcPort = srcPort
    return b
}
func (b *builder) WithDestAddr(destAddr string) *builder {
    b.msg.Header.DestAddr = destAddr
    return b
}
func (b *builder) WithDestPort(destPort uint64) *builder {
    b.msg.Header.DestPort = destPort
    return b
}
func (b *builder) WithHeaderItem(key, value string) *builder {
  // 保证map只初始化一次   => 这里是个知识点
    b.once.Do(func() {
        b.msg.Header.Items = make(map[string]string)
    })
    b.msg.Header.Items[key] = value
    return b
}
func (b *builder) WithBodyItem(record string) *builder {
    b.msg.Body.Items = append(b.msg.Body.Items, record)
    return b
}
// 创建Message对象，在最后一步调用
func (b *builder) Build() *Message {
    return b.msg
}
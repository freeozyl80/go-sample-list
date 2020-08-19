package msg

// 还可以这样循环着玩
type Prototype interface {
    Clone() Prototype
}

type Message struct {
    Header *Header
    Body   *Body
}

type Header struct {
    SrcAddr  string
    SrcPort  uint64
    DestAddr string
    DestPort uint64
    Items    map[string]string
}

type Body struct {
    Items []string
}

func (m *Message) Clone() Prototype {
    msg := *m
    return &msg
}
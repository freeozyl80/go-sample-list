package event

type Type uint8

const (
    Start Type = iota
    End
)

// 事件抽象接口
type Event interface {
    EventType() Type
    Content() string
}
// 开始事件，实现了Event接口
type StartEvent struct{
    content string
}
// 结束事件，实现了Event接口
type EndEvent struct{
    content string
}

func(e * StartEvent) EventType() Type{
	return 0
}

func(e * EndEvent) EventType() Type{
	return 1
}

func(e * StartEvent) Content() string{
	return e.content
}

func(e * EndEvent) Content() string{
	return e.content
}
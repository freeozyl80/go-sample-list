package event

type Factory struct{}
// 更具事件类型创建具体事件
func (e *Factory) Create(etype Type) Event {
    switch etype {
    case Start:
        return &StartEvent{
            content: "this is start event",
        }
    case End:
        return &EndEvent{
            content: "this is end event",
        }
    default:
        return nil
    }
}
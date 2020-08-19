package plugin

import (
	"reflect"
)
// input插件名称与类型的映射关系，主要用于通过反射创建input对象
var inputNames = make(map[string]reflect.Type)
// Hello input插件，接收“Hello World”消息
type HelloInput struct {}

func (h *HelloInput) Receive() string {
    return "Hello World"
}
// 初始化input插件映射关系表
func init() {
    inputNames["hello"] = reflect.TypeOf(HelloInput{})
}

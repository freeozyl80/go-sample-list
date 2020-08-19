package plugin

import (
	"fmt"
	"reflect"
)
// output插件名称与类型的映射关系，主要用于通过反射创建output对象
var outputNames = make(map[string]reflect.Type)
// Console output插件，将消息输出到控制台上
type ConsoleOutput struct {}

func (c *ConsoleOutput) Send(msg string) {
    fmt.Println(msg)
}
// 初始化output插件映射关系表
func init() {
    outputNames["console"] = reflect.TypeOf(ConsoleOutput{})
}
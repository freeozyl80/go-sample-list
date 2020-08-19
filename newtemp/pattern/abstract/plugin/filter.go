package plugin

import (
	"reflect"
	"strings"
)

// filter插件名称与类型的映射关系，主要用于通过反射创建filter对象
var filterNames = make(map[string]reflect.Type)
// Upper filter插件，将消息全部字母转成大写
type UpperFilter struct {}

func (u *UpperFilter) Process(msg string) string {
    return strings.ToUpper(msg)
}
// 初始化filter插件映射关系表
func init() {
    filterNames["upper"] = reflect.TypeOf(UpperFilter{})
}
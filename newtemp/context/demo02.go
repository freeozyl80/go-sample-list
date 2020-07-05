package main

import (
    "fmt"
    "time"
    "context"
)

var key string = "name"

// 使用通过context向goroutinue传递值
func watch(ctx context.Context) {
    for {
        select{
        case <- ctx.Done():
            fmt.Println(ctx.Value(key), "退出，停止了。。。")
            return
        default:
            fmt.Println(ctx.Value(key), "运行中...")
            time.Sleep(2 * time.Second)
        }
    }
} 

func main() {
    ctx, cancel := context.WithCancel(context.Background())
  	// 给ctx绑定键值，传递给goroutine
    valuectx := context.WithValue(ctx, key, "【监控1】")
  	// 启动goroutine
    go watch(valuectx)

    time.Sleep(time.Second * 10)
    fmt.Println("该结束了。。。")
  	// 运行结束函数
    cancel()
    time.Sleep(time.Second * 3)
    fmt.Println("真的结束了。。")
}

// context.Background() 返回一个空的 Context，这个空的 Context 一般用于整个 Context 树的根节点。然后我们使用 context.WithCancel(parent) 函数，创建一个可取消的子 Context，然后当作参数传给 goroutine 使用，这样就可以使用这个子 Context 跟踪这个 goroutine。
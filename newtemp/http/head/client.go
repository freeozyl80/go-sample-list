package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
    fmt.Println("----------访问一个正常的 URL ----------")
    MyHead("https://www.baidu.com")
    // 结果：返回的状态码是： 200

    fmt.Println("----------访问一个正常的 URL ----------")
    // 可以将之前测试 Do 用的服务器端启动，因为这个服务器端有限制，
    // 可以通过返回的 code 来直到访问的状况。
    MyHead("http://localhost:8080")
    // 结果： 返回的状态码是： 400
}

func MyHead(url string) {
    resp, err := http.Head(url)
    ErrPrint(err)

    defer resp.Body.Close()

    bytes, err := ioutil.ReadAll(resp.Body)
    ErrPrint(err)

    fmt.Printf("%q\n", bytes)
    fmt.Printf("返回的状态码是： %d\n", resp.StatusCode)
}

func ErrPrint(err error) {
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
}


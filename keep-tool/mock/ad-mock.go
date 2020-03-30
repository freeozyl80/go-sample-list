package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("正在启动server => port:9090\n")
	http.HandleFunc("/mock/ad/", mockData)   //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func mockData(w http.ResponseWriter, r *http.Request) {
	var materailRegex = regexp.MustCompile(`^material`)
	var cmsproxyRegex = regexp.MustCompile(`^cmsproxy`)

	pth := r.URL.Path
	str := strings.Split(pth, "mock/ad/")[1]

	switch {
	case materailRegex.MatchString(str):
		fmt.Printf("进入ad")
		file := strings.Split(str, "material")[1]
		file = "./material-json" + file + ".json"

		var v interface{}
		JsonParse := NewJsonStruct()

		JsonParse.Load(file, &v)

		data, _ := json.Marshal(v)

		w.WriteHeader(200)

		fmt.Fprintf(w, "%+v", string(data))
	case cmsproxyRegex.MatchString(str):
		fmt.Printf("进入test")
		for _, cookie := range r.Cookies() {
			fmt.Println(w, cookie.Name)
		}
	}
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, v)

	if err != nil {
		return
	}
}

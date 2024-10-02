package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

/*
	我这样理解: 利用各个语言对SocketApi的封装实现我们可以实现数据的传输与接收, 而Socket利用像TCP, UDP这样的传输层协议实现上层传输以及接收数据
	继续说我对应用层HTTP协议的理解, 简单来说就是约定了客户端和服务器之间的一种请求响应关系,
	通过协议约定好的请求格式, 以及响应格式底层通过socket就可以实现资源request与response,
	go语言中net封装了Socket对象的基本功能而在net的基础上实现了http这个标准库则实现了HTTP协议约定好的请求响应格式让我们专注于请求体以及响应体的实现而并不需要在net库中Socket对象基础上构建这种格式
*/

// /index的处理函数
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	indexData, _ := ioutil.ReadFile("index.html")
	fmt.Fprint(w, string(indexData))
}

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/w", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello")

	})

	// 服务监听地址, 以及协议种类
	http.ListenAndServe(":8080", mux)
}

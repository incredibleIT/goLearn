package main

import (
	"fmt"
	"net"
)

/*
	cs : client-server 客户端-服务端
	bs : browser-server 浏览器-服务端  固化了客户端为浏览器
	web应用程序 : 针对浏览器作为客户端发送的请求做出的处理应用
	http协议: 应用层协议, 基于TCP/IP协议, 短连接, 基于请求响应(浏览器发, 服务器响应), 无状态保存
		请求协议格式(request) : 限定浏览器给服务器发送数据格式
		响应协议格式(response): 限定服务器给浏览器发送数据格式
*/

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("err", err)
	}
	defer listener.Close()

	conn, _ := listener.Accept()
	fmt.Println(conn)

	data := make([]byte, 1024)
	n, _ := conn.Read(data)

	fmt.Println(string(data[:n]))

	conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/html; charset=utf-8\r\n\r\nYangYang"))

}

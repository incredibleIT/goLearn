package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Addr string
	Chan chan []byte
}

type Msg struct {
	Addr    string
	Content string
}

// 广播通道
var broadcast = make(chan Msg)

// 广播的在线客户端
var onlineClient = make(map[string]Client)

func messageManager() {
	for {
		broadcastMsg := <-broadcast
		for _, client := range onlineClient {
			broadcastMsgByte, _ := json.Marshal(broadcastMsg)
			client.Chan <- broadcastMsgByte
		}
	}
}

func read(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		broadcast <- Msg{Addr: conn.RemoteAddr().String(), Content: strings.TrimSpace(string(data[:n]))}

		fmt.Printf("[%s] : %s\n", conn.RemoteAddr().String(), strings.TrimSpace(string(data[:n])))
	}
}

func write(conn net.Conn, client Client) {
	for {
		data := <-client.Chan
		conn.Write(data)
	}
}

func ConnHandler(conn net.Conn) {
	s := conn.RemoteAddr().String()
	fmt.Printf("[%s]成功连接\n", s)

	client := Client{Addr: s, Chan: make(chan []byte)}
	onlineClient[s] = client

	broadcast <- Msg{Addr: "系统消息", Content: "[" + s + "]" + " 上线啦"}
	go read(conn)
	go write(conn, client)
}

// 服务端
func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	go messageManager()

	for {
		conn, _ := listener.Accept()
		go ConnHandler(conn)
	}

}

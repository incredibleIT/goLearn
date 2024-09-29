package main

import (
	"fmt"
	"net"
	"strings"
)

/*
	各个高级语言的网络编程库通常是对底层 socket API 的封装，简化了网络通信的复杂性。
	通过这些封装，开发者可以更方便地建立连接、发送和接收数据，而无需直接处理 socket 的底层细节。
	这样可以提高开发效率和代码可读性。


	Socket API 通常是用 C 语言实现的，因为它最早是在 Unix 系统上开发的。
	不过，许多其他编程语言（如 Python、Java、Go 等）都提供了对 Socket API 的封装，
	使得在这些语言中进行网络编程变得更加简单和直观。
*/

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("err: ", err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("err: ", err)
		}

		fmt.Println("conn: ", conn)

		for {
			b := make([]byte, 1024)
			n, err := conn.Read(b)

			if err != nil {
				fmt.Println("err: ", err)
			}
			clientData := string(b[:n])
			if clientData == "exist" {
				break
			}
			// server端拿到数据
			fmt.Println(clientData)

			// 做转大写处理
			clientData = strings.ToUpper(clientData)

			// 传回
			n, _ = conn.Write([]byte(clientData))
		}
	}
}

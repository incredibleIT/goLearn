package main

import (
	"fmt"
	"net"
)

// 一个连接服务端的客户端
func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	fmt.Println(conn)

	if err != nil {
		fmt.Println("err: ", err)
	}

	for {

		var k string

		fmt.Scan(&k)

		n, _ := conn.Write([]byte(k))
		fmt.Println("n :", n)

		if k == "exist" {
			break
		}

		// 接收大写
		serverData := make([]byte, 1024)
		n, _ = conn.Read(serverData)

		fmt.Println("serverData:", string(serverData[:n]))
	}

}

package main

import (
	"fmt"
	"net"
	"strings"
)

func getDo(conn net.Conn) {
	for {
		var clientData = make([]byte, 1024)
		n, _ := conn.Read(clientData)
		data := strings.ToUpper(string(clientData[:n]))

		fmt.Println("接收到字节数 : ", n, "接收到数据 : ", string(clientData[:n]))

		n, _ = conn.Write([]byte(data))
		fmt.Println("传输字节数 : ", n)
	}
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		go getDo(conn)

	}

}

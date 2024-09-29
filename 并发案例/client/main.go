package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	fmt.Println("conn : ", conn)

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("请输入 >>")
		clientData, _ := reader.ReadString('\n')

		n, _ := conn.Write([]byte(clientData))

		fmt.Println("传到客户端字节数 : ", n)

		ServerData := make([]byte, 1024)
		n, _ = conn.Read(ServerData)

		fmt.Println("接收字节数 : ", n)

		fmt.Println(string(ServerData[:n]))
	}
}

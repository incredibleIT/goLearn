package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Msg struct {
	Addr    string
	Content string
}

func read(conn net.Conn) {
	for {
		var data = make([]byte, 1024)
		msg := Msg{}
		n, _ := conn.Read(data)
		json.Unmarshal(data[:n], &msg)

		fmt.Printf("[%s] : %s\n", msg.Addr, msg.Content)
	}
}

func write(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadBytes('\n')

		conn.Write(data)
	}

}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go read(conn)
	write(conn)

}

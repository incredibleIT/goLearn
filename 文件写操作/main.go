package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 方式一
func write01(file *os.File) {

	s := "hello\n"

	n, _ := file.Write([]byte(s)) // file直接调Write()写的是字节串

	num, _ := file.WriteString(s)

	fmt.Println(n)
	fmt.Println(num)

}

// 方式二
func write02(file *os.File) {

	writer := bufio.NewWriter(file)

	// writer.Write()
	// writer.WriteString()

	writer.WriteString("world\n") // 这一步不会书写, 先缓存达到一定量再写

	writer.Flush() // 这一步才会书写, 将缓存部分全部书写

}

// 方式三
func write03() {

	err := ioutil.WriteFile("文本1", []byte("hello world\n"), 0666)

	if err != nil {
		fmt.Println(err)
	}

}

func main() {

	file, _ := os.OpenFile("文本new", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	defer file.Close()

	//write01(file)

	//write02(file)

	write03()

}

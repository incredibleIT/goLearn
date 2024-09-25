package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// *os.File类型的对象中Read方法实现按字节串的读取
func read01(file *os.File) {
	b := make([]byte, 3)
	n, _ := file.Read(b)
	fmt.Print(string(b), n)

}

// bufio 包装了 io.Reader 和 io.Writer 接口, 允许我们以缓冲的方式进行读取和写入操作，特别适合处理大数据量或频繁的小数据操作。
func read02(file *os.File) {

	// 通过bufio的NewReader方法进行reader对象的构建
	reader := bufio.NewReader(file)

	// ReadString() 参数为截至符, 返回读出的字符串和err
	//s, _ := reader.ReadString('\n')
	//
	//fmt.Print(s)

	for {
		//line, err := reader.ReadString('\n')
		//
		//fmt.Print(line)
		//
		//if err == io.EOF {
		//	break
		//}

		// ReadBytes(delim) 方法会读出字节串, 可通过强转转为字符串 参数为截至符, 返回一行数据的字节串和err
		line, err := reader.ReadBytes('\n')

		fmt.Print(line)
		fmt.Print(string(line))

		if err == io.EOF {
			break
		}

	}

}

func read03() {
	file, err := ioutil.ReadFile("文本")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func main() {

	// os.Open()来打开文件, 反回句柄, 和err
	file, _ := os.Open("文本")

	//read01(file)

	read02(file)

	//read03()
}

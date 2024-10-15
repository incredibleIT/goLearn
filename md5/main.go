package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md := md5.New()

	md.Write([]byte("hello world"))

	// 加密
	sum := md.Sum(nil) // 128bit 传回的使byte切片, 八位表示一个字节

	fmt.Println(sum) // 128 / 8 = 16字节  那么len(sum) == 16

	toString := hex.EncodeToString(sum) // 可读 // 将字节切片转换为十六进制字符串    1B = 8b   16个字节  1个十六进制位表示半个字节, 那么1字节由2个十六进制位表示  16字节由32个十六进制位表示

	fmt.Println(len(toString))

}

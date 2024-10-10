package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	logfile, err := os.Create("./log" + ".txt")

	if err != nil {
		fmt.Println(err)
	}

	loger := log.New(logfile, "test_", log.LstdFlags|log.Lshortfile)

	loger.Println("这是错误1111")

	loger.Output(2, "打印错误信息")

	loger.Print("这是错误222222")

	fmt.Println(loger.Flags())

	fmt.Println(loger.Prefix())

}

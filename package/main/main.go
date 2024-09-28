package main

/*
	在本地包管理中对于引入的包, go首先会去goroot 拼接src路径中找对应的包
	如果goroot没有对应的包会去gopath然后拼接src路径来寻找自定义的包, 找到后导入
	对于网络中的包可执行go get 包路径 来进行下载 然后go会将其保存至go path src路径下然后导入

	函数首字母大写代表公共, 小写代表私有
*/
import (
	"awesomeProject/package/utils"
	"fmt"
)

func main() {

	fmt.Println(utils.Add(1, 2))

}

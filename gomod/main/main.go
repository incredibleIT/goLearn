package main

/*
	go mod init 来初始化一个mod文件来用于管理依赖
	go mod tidy 用于清理和整理 Go Modules 的依赖关系,删除未使用的依赖,添加缺失依, 更新go.sum
	go mod vendor 用于将依赖导入项目本地的vendor目录
*/

import (
	"fmt"
	"github.com/jinzhu/now"
	"gomod/utils"
)

func main() {

	fmt.Println(utils.Add(1, 2))

	fmt.Println(now.BeginningOfMinute())
}

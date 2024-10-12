package confTest

import (
	"fmt"
	"testing"
	"viperdemo/utils"
)

func TestGet(test *testing.T) {
	// 读取
	dbviper := utils.GetConfig("mysql")

	fmt.Println(dbviper.GetInt("blog.port"))

}

// 今天学习了: 1. 编写单元测试, 基准测试
//			  2. rune数据类型: 为了解决多个字节存储字符问题
//			  3.原生库runtime获取当前项目根目录
//			  4.viper加载以及读取配置文件

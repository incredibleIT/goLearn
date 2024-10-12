package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

//var projectParentPath = path.Dir(getParentPath()+"/../") + "/"
//
//func getParentPath() string {
//	_, fileName, _, _ := runtime.Caller(0) // 当前路径
//
//	return path.Dir(fileName)
//}
//
//// *viper.Viper
//func GetConfig(file string) *viper.Viper {
//	fmt.Println("projectParentPath: ", projectParentPath)
//	configer := viper.New()
//	fmt.Println(projectParentPath)
//	confpath := projectParentPath + "configs/"
//
//	configer.AddConfigPath(confpath)
//	configer.SetConfigName(file)
//	configer.SetConfigType("yaml")
//
//	configFile := confpath + file + ".yaml"
//
//	if err := configer.ReadInConfig(); err != nil {
//		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
//			panic(fmt.Errorf("未找到配置文件%s", configFile))
//		} else {
//			panic(fmt.Errorf("解析配置文件%s出错:%s", configFile, err))
//		}
//	}
//
//	fmt.Println(configFile)
//
//	return configer
//
//}

var projectRootPath = path.Dir(getCurrentPath()+"/../") + "/"

// 得到项目配置文件目录
func GetConfig(file string) *viper.Viper {
	configPath := projectRootPath + "configs/"
	configer := viper.New()

	configer.AddConfigPath(configPath)
	configer.SetConfigName(file)
	configer.SetConfigType("yaml")

	//fmt.Println(configPath + file + ".yaml")
	if err := configer.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("配置文件%s未找到", configPath+file+".yaml"))
		} else {
			panic(fmt.Errorf("解析配置文件%s出错:%s", configPath+file+".yaml", err))
		}
	}

	return configer
}

// 得到当前位置
func getCurrentPath() string {
	_, filePath, _, _ := runtime.Caller(0)

	return path.Dir(filePath) // viper/utils/
}

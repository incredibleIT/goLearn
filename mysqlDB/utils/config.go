package creatConnection

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var RootPath = path.Dir(getCurrentPath()+"/../") + "/" //     mysqlDB/

func getCurrentPath() string {
	_, currentPath, _, _ := runtime.Caller(0)

	return path.Dir(currentPath) //    mysqlDB/utils/
}

func GetConfig(fileName string) *viper.Viper {
	configer := viper.New()
	configPath := RootPath + "configs/"
	configer.AddConfigPath(configPath)
	configer.SetConfigName(fileName)
	configer.SetConfigType("yaml")
	filePath := configPath + fileName + ".yaml"

	if err := configer.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("配置文件%s未找到", filePath))
		} else {
			panic(fmt.Errorf("配置文件%s解析错误: %s", filePath, err))
		}
	}

	return configer

}

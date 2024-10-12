package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var ParentPath = path.Dir(GetCurrentPath()+"/../") + "/" // logrus.../

func GetConfig(confFileName string) *viper.Viper {

	Viper := viper.New()
	configPath := ParentPath + "configs/"
	Viper.AddConfigPath(configPath)
	Viper.SetConfigName(confFileName)
	Viper.SetConfigType("yaml")

	realFilePath := configPath + confFileName + ".yaml"

	if err := Viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("配置文件%s不存在", realFilePath))
		} else {
			panic(fmt.Errorf("配置文件%s解析出错:%s", realFilePath, err))
		}
	}

	return Viper

}

// 获取当前目录
func GetCurrentPath() string {
	_, file, _, _ := runtime.Caller(0)

	return path.Dir(file) // logrus + ro.../utils/
}

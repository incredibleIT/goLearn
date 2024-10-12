package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var Logrus *logrus.Logger

func InitLog(configFile string) {

	Viper := GetConfig(configFile)

	Logrus = logrus.New()

	switch strings.ToLower(Viper.GetString("level")) {
	case "debug":
		Logrus.SetLevel(logrus.DebugLevel)
	case "info":
		Logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		Logrus.SetLevel(logrus.WarnLevel)
	case "error":
		Logrus.SetLevel(logrus.ErrorLevel)
	case "panic":
		Logrus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("错误的log等级%s", Viper.GetString("level")))
	}

	// 配置日志格式
	Logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	// log记录文件位置
	logPath := ParentPath + Viper.GetString("filePath")
	fmt.Println(logPath)

	fout, err := rotatelogs.New(
		logPath+".%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithRotationTime(1*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)

	if err != nil {
		panic(err)
	}

	Logrus.SetOutput(fout)
	Logrus.SetReportCaller(true)

}

package loggerTest

import (
	logger "logger/utils"
	"testing"
)

func TestLogger(t *testing.T) {
	logger.InitLog("log")

	logger.Logrus.Debug("debug......................log")
	logger.Logrus.Info("info......................log")
	logger.Logrus.Warn("warn......................log")
	logger.Logrus.Error("error......................log")
	//logger.Logrus.Panic("panic......................log")

}

// 今天学习了: 1.复习viper解析配置文件
//			 2. logrus日志工具, 多个日志等级debug, info, warn, error, panic, 设置logrus日志格式
//			 3.rotatelogs日期轮询日志, 通过New()rotatelogs对象传入参数进行日志输出路径配置, 轮询时间, 以及存留时间配置

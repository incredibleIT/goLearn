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

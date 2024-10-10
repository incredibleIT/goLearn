package main

import (
	"go.uber.org/zap"
	"net/http"
)

var zaper *zap.Logger

func main() {

	// logger
	initZap()
	defer zaper.Sync()
	GetHttp("https://www.baidu.com/")

}

func initZap() {
	zaper, _ = zap.NewProduction()
}

func GetHttp(url string) {

	res, err := http.Get(url)

	if err != nil {
		zaper.Error(
			"没有匹配的url",
			zap.String("url", url),
			zap.Error(err))
	} else {
		zaper.Info("成功啦",
			zap.String("url", url),
			zap.String("status", res.Status))
	}

}

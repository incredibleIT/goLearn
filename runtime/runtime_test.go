package runtime

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func fa() (string, int) {
	_, fileName, line, _ := runtime.Caller(2)

	return fileName, line
}

func fb() (string, int) {
	return fa()
}

func TestCaller(t *testing.T) {
	fileName, line := fb()

	fmt.Println(fileName, line)
	fmt.Println(path.Dir(fileName))
	fmt.Println(path.Dir(path.Dir(fileName) + "/../")) // 项目根目录
}

package stringtest

import (
	"testing"
	tostring "unitTest"
)

func TestIntToString(t *testing.T) {
	i := 12
	s := tostring.IntToString(i)

	if s == "12" {
		t.Fail()
	}
}

// go test -v ./ -run=TestIntToString -count=1
// go test: 命令  -v: 希望将输出语句输出  -run=(正则): 想要执行的测试方法名  -count: 希望编译后复用次数

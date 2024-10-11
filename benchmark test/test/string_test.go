package string_test

import (
	tostring2 "benchmarkTest"
	"testing"
)

// 单元测试, 接收*testing.T类型参数
func TestIntToString(test *testing.T) {
	i := 133
	s := tostring2.IntToString(i)

	if s != "133" {
		test.Fail()
	}
}

// 基准测试
// Benchmark开头, 接受参数*testing.B类型
// 动用b.N次数的循环来进行测试
func BenchmarkIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tostring2.IntToString(i)
	}
}

//  go test ./test -run=^$ -count=1 -bench=BenchmarkIntToString
//  go test ./test -run=^$ -count=1 -benchmem -bench=BenchmarkIntToString
//  go test ./test -run=^$ -count=1 -benchmem -bench=BenchmarkIntToString -memprofile=mem -cpuprofile=cpu

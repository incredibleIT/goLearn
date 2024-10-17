package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type s struct {
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

// 时间序列化与反序列化
func main() {

	/*
		当使用"encoding/json"包来处理JSON时, go的time.Time类型会自动与劣化为RFC3339格式
		同时反序列化也会自动解析
	*/

	s1 := s{
		Name: "yangyang",
		Time: time.Now()}

	s1Json, _ := json.Marshal(s1)

	fmt.Println(string(s1Json))

	// 反序列化
	if err := json.Unmarshal(s1Json, &s1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s1)
		fmt.Println(s1.Time.Format("2006-01-02 15:04:05"))

	}
}

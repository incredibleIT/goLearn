package main

import (
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

	//s1 := s{
	//	Name: "yangyang",
	//	Time: time.Now()}
	//
	//s1Json, _ := json.Marshal(s1)
	//
	//fmt.Println(string(s1Json))
	//
	//// 反序列化
	//if err := json.Unmarshal(s1Json, &s1); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(s1)
	//	fmt.Println(s1.Time.Format("2006-01-02 15:04:05"))
	//
	//}

	//ticker := time.NewTicker(2 * time.Second)
	//
	//<-ticker.C
	//
	//ticker.Reset(time.Duration(3) * time.Second)
	//
	//<-ticker.C
	//
	//fmt.Println("heloo")

	ScheduleTask(func() {
		fmt.Println("hello")
		time.Sleep(1 * time.Second)
	}, 15, 34, 0)

}

var loc, _ = time.LoadLocation("Asia/Shanghai")

func ScheduleTask(work func(), hour int, minute int, second int) error {
	var oneDay int64 = 86400
	ticker := time.NewTicker(1)

	// 总体上就是一个无尽的for循环, 每次for循环判断 当前时间 与 任务执行时间 之间相差的秒数, 然后阻塞相应的描述之后执行任务重新循环
	for {
		now := time.Now()                                                                    // 返回一个time.Time类型
		point := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, loc) // 构建定时任务需要执行的时间

		// 要计算当前时间与任务执行时间相差多秒, differ
		differ := (point.Unix()%oneDay + oneDay - now.Unix()%oneDay) % oneDay // 计算相差时间

		fmt.Println(differ)

		if differ == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		// 程序要阻塞相应时间
		ticker.Reset(time.Duration(differ) * time.Second)
		fmt.Println("reseted")
		<-ticker.C

		fmt.Println("work")
		work()
	}
}

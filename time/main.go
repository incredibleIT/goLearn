package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	// time标准库, 时间处理

	// 获取当前时间
	now := time.Now()
	fmt.Println(now) //2024-10-16 00:29:22.6664279 +0800 CST m=+0.002721601

	// 时间格式化输出, 将时间对象转换为字符串
	fmt.Println(now.Format("2006-01-02 15:04:05.000")) // 2024-10-16 00:31:15.993

	// 将字符串转化为时间对象
	layOut := "2006-01-02 15:04:05.000" // 时间格式化模板
	ts := "2024-10-16 00:31:15.993"
	tm, _ := time.Parse(layOut, ts)
	fmt.Println(tm)
	// 若此时想指定试时区
	loc1, _ := time.LoadLocation("Asia/Shanghai")
	tm1, _ := time.ParseInLocation(layOut, ts, loc1)
	fmt.Println(tm1)

	// 时区  +0800是东八区: 北京时间
	// 例如如果想获取美国的时间
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.In(loc)) // 获得当前时区的具体时间

	// 可以用time.Date()来指定一个特定时区的时间
	t := time.Date(2024, 1, 1, 22, 30, 50, 0, loc)
	fmt.Println(t)

	// 时间运算
	fmt.Println(now.Add(1 * time.Hour))
	fmt.Println(now.Add(1 * time.Second))
	fmt.Println(now.Add(2 * time.Millisecond))
	fmt.Println(now.Add(3 * time.Microsecond)) // 微秒
	fmt.Println(now.Add(4 * time.Nanosecond))  // 纳秒

	fmt.Println(now.AddDate(0, 1, 1))

	// 减法
	tm2 := now.Add(1 * time.Hour)
	fmt.Println(now.Sub(tm2))

	// 时间的比较
	now.Equal(tm2)  // ==
	now.After(tm2)  // >
	now.Before(tm2) // <

	// 时间戳   表示自 1970 年 1 月 1 日 00:00:00 UTC 以来经过的秒数。  比较时间大小也可以用时间戳来比
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli()) // 毫秒数
	fmt.Println(now.UnixMicro()) // 微秒数
	fmt.Println(now.UnixNano())  // 纳秒数

	// Duration(时间间隔), 是time包中一个非常重要的概念, 表示的是两个时间点之间的时间长度

	// time.Duration类型是一个表示纳秒的整数类型, 可以用来表示时间间隔.可以使用它进行加减法运算

	// 创建time.Duration对象
	// 1. 使用time.Second, time.Hour, time.Minute 等常量创建
	dur1 := time.Second * 1
	dur2 := time.Minute * 2
	dur3 := time.Hour * 3
	fmt.Println(dur1, dur2, dur3, reflect.TypeOf(dur1), reflect.TypeOf(dur2), reflect.TypeOf(dur3))

}

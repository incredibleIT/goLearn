package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type stu1 struct {
}

func (s stu1) Run() {
	fmt.Println("stu1 run")
}

type stu2 struct {
}

func (s stu2) Run() {
	fmt.Println("stu2 run")
}

func (s stu2) Keep() {
	fmt.Println("stu2 keep")
}

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func main() {
	i := 0
	c := newWithSeconds()

	_, err := c.AddFunc("*/5 * * * * ?", func() {
		fmt.Println("cron running: ", i)
		i++
	})

	if err != nil {
		panic(err)
	}

	_, err = c.AddJob("*/2 * * * * ?", stu1{})
	_, err = c.AddJob("*/4 * * * * ?", stu2{})

	c.Start()

	select {}

}

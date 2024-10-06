package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model

	Name string

	Age int
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	student01 := Student{Name: "yangyang", Age: 20}

	// insert一条记录
	db.Create(&student01)
	fmt.Println(student01)

	// insert批量插入
	studentList := []Student{{Name: "yangyang02", Age: 1}, {Name: "yangyang03", Age: 2}, {Name: "yangyang04", Age: 3}}
	db.Create(&studentList)

	// 指定特定的字段插入
	db.Select("Age").Create(&student01)
	// 指定特定字段之外的字段插入
	db.Select("Age").Create(&student01)

	// 可通过map结构进行插入, 但自动填充的字段都会失效
	db.Model(&Student{}).Create(map[string]interface{}{"name": "yangyang05"})

}

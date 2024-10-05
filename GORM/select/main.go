package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday *time.Time
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var user User
	//
	//// 没有找到记录时，它会返回 ErrRecordNotFound 错误
	//// 查询第一条
	//db.First(&user)
	//// 查询一条, 没有指定排序策略
	//db.Take(&user)
	//// 查询最后一条
	//db.Last(&user)

	//fmt.Println(user.ID)
	//fmt.Println(user.Name)
	//fmt.Println(user.Age)
	//fmt.Println(user.Birthday)

	/*
		通过主键检索
		db.First(&user, 10)
		// SELECT * FROM users WHERE id = 10;

		db.First(&user, "10")
		// SELECT * FROM users WHERE id = 10;

		db.Find(&users, []int{1,2,3})
		// SELECT * FROM users WHERE id IN (1,2,3);
	*/

	//db.First(&user, 1)
	//
	//fmt.Println(user.ID)
	//fmt.Println(user.Name)
	//fmt.Println(user.Age)
	//fmt.Println(user.Birthday)

	// 检索全部对象

	userList := make([]User, 10)
	db.Find(&userList)
	for _, user := range userList {
		fmt.Println(user.Name)
	}

	// 条件查询

	// string条件, 字符串拼接
	db.Where("name", "yangyang").First(&user)
	db.Where("name IN", []string{"yangyang", "yangyang2", "yangyang3"}).First(&user)
	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&user)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	db.Where("updated_at > ?", "lastWeek").Find(&user)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", "lastWeek", "today").Find(&user)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	fmt.Println(user.Name)

	// Struct条件, Map条件, 主键切片
	db.Where(&User{Name: "yangyang"}).First(&user)

	db.Where(map[string]interface{}{"name": "yangyang"}).First(&user)

	db.Where([]int64{1, 2, 3}).First(&user)

	// 当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件, 但可用map构建他们

	// 内联条件
	// SELECT * FROM users WHERE id = 23;
	// 根据主键获取记录，如果是非整型主键
	db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// Plain SQL
	db.Find(&userList, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&userList, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&userList, User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&userList, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;

}

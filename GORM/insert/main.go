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

// 钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	fmt.Println("before create user")

	return
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	if err != nil {
		panic(err)
	}

	//user := User{Name: "yang", Age: 18, Birthday: time.Now()}
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	panic(err)
	//}

	// 1. 插入数据
	//result := db.Create(&user)
	//
	//fmt.Println(user.ID)             // 返回插入数据的主键
	//fmt.Println(result.Error)        // 返回 error
	//fmt.Println(result.RowsAffected) // 返回插入记录的条数

	// 2. 用指定的字段创建记录
	//db.Select("Age").Create(&user) // 创建记录并更新给出的字段

	//db.Omit("Age").Create(&user) // 创建记录并更新未给出的字段

	// 3. 批量插入
	//userList := []User{{Name: "y1"}, {Name: "y2"}, {Name: "y3"}, {Name: "y4"}, {Name: "y5"}}

	//db.Create(&userList)
	//
	//for _, u := range userList {
	//	fmt.Println(u.ID)
	//}

	// CreateInBatches() 可以指定每批插入的数量
	//db.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(userList, 3) // 跳过购自函数
	//
	//for _, user := range userList {
	//	fmt.Println(user.ID)
	//}

	// db.Model() 是 GORM 中用于指定模型的一个方法，通常用于更新、删除或进行复杂查询时。它允许你明确指定要操作的数据库表和模型。
	// 允许map创建记录, 不过自动填充都没有了
	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "yyy",
	//})

	//user := User{
	//	Name:     "yangyangyang",
	//	Age:      12,
	//	Birthday: nil,
	//	Person: Person{
	//		Number: 1111,
	//	},
	//}
	//db.Create(&user)
}

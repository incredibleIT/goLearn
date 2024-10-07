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

	// save方法  会将所有值保存一遍, 包括零值
	var student Student
	db.First(&student)
	fmt.Println(student)
	student.Age = 11
	db.Save(&student)

	// update方法

	// (1) 更新单列, 需要指定条件, 使用Model方法并且主键有值会被构建条件

	// 条件更新
	db.Debug().Model(&Student{}).Where("id = ?", 3).Update("name", "yang")
	// Model值更新
	db.Debug().Model(&student).Update("name", "yang1")
	// 根据条件, Model值更新
	db.Debug().Model(&student).Where("name = ?", "yang1").Update("age", 111)

	// (2) 更新多列 Updates方法

	// struct方式更新, 不会更新为零值的字段
	db.Debug().Model(&student).Updates(Student{Name: "yang3", Age: 0})

	// map格式更新, 会更新零值
	db.Debug().Model(&student).Updates(map[string]interface{}{"name": "yang4", "age": 0})

	// 可使用select 和 omit 来选定或规避更新字段

	db.Debug().Model(&student).Select("name").Updates(Student{Name: "yang4", Age: 111111}) // 不会更新age, 因为Select选定了name字段

	db.Debug().Model(&student).Omit("name").Updates(Student{Name: "yang5", Age: 222222}) // 会更新除了name字段其他的所有字段

	// 通过select可以结合struct结构进行零值的更新
	db.Debug().Model(&student).Select("age").Updates(Student{Name: "yang6", Age: 0})

	// 批量更新  如果未通过Model来指定记录的主键, 会执行批量更新
	db.Debug().Where("name = ?", "yang6").Updates(Student{Name: "yang7", Age: 1111})

	// 如果在没有任何条件的情况下执行批量更新，默认情况下，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	//db.Model(&Student{}).Update("name", "yang")  报错
	db.Debug().Model(&Student{}).Where("1 = 1").Update("name", "yang")

	// 获取影响行数
	result := db.Debug().Model(&student).Update("name", "yyyy")
	fmt.Println(result.RowsAffected) // 影响行数
	fmt.Println(result.Error)        // 更新的错误

}

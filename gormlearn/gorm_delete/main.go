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

// 带有gorm.deletedat 字段会自动被软删除处理
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 当删除时需要实体指定id值, 不然触发批量删除
	var student = Student{Model: gorm.Model{ID: 2}}

	db.Debug().Delete(&student)

	// delete允许通过主键来进行检索删除
	db.Debug().Delete(&Student{}, 10)
	db.Debug().Delete(&Student{}, []int{6, 7, 8, 9, 10})

	// 条件指定的值不包括主属性, 那么会批量删除
	db.Debug().Where("name = ?", "yangyangyang").Delete(&Student{})
	db.Debug().Delete(&Student{}, "name = ?", "yangyangyang")

	// 如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	//db.Debug().Delete(&Student{})  报错

	db.Debug().Where("1=1").Delete(&Student{}) // 全删, 要有条件

	// 通过Unscoped可以查询所有软删除的记录
	var students []Student
	db.Debug().Unscoped().Find(&students)
	for _, student := range students {
		fmt.Println(student)
	}

	// 通过Unscoped实现永久删除
	db.Debug().Delete(&student)
}

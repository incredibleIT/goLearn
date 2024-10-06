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

	student := Student{}
	// 查询以id升序排列的第一条记录
	db.First(&student)
	fmt.Println(student.Name, student.Age)

	// 查询自然排序的一条记录
	db.Take(&student)
	fmt.Println(student.Name, student.Age)

	// 查询以id升序排列的最后一条记录
	db.Last(&student)
	fmt.Println(student.Name, student.Age)

	// 查询所有记录
	var studentList []Student
	db.Find(&studentList)
	fmt.Println(studentList)

	// 条件查询
	//(1) string条件拼接
	db.Where("id = ?", 2).Find(&student)
	fmt.Println(student)

	db.Where("name in ?", []string{"yangyang01", "yangyang", "yangyang02"}).Find(&studentList)
	fmt.Println(studentList)

	db.Where("name like ?", "y%").Find(&studentList)
	fmt.Println(studentList)

	// (2) struct条件, 如果某一字段为其零值不会作为条件去查询
	db.Debug().Where(Student{Name: "yangyang"}).Find(&studentList)
	fmt.Println(studentList)

	db.Debug().Where(Student{Name: "yangyang02", Age: 11}).Find(&studentList)
	fmt.Println(studentList)

	// 不会拼接零值
	db.Debug().Where(Student{Name: "yangyang01", Age: 0}).Find(&studentList)
	fmt.Println(studentList)

	// (3)map条件, 查零值可以用map拼
	db.Debug().Where(map[string]interface{}{"name": "yangyang02"}).Find(&studentList)
	fmt.Println(studentList)

	// 可查零值
	db.Debug().Where(map[string]interface{}{"name": "yangyang01", "Age": 0}).Find(&studentList)
	fmt.Println(studentList)

	// Not语句
	db.Debug().Not("name = ?", "yangyang02").Find(&studentList)
	db.Debug().Not(map[string]interface{}{"name": "yangyang02"}).Find(&studentList)

	// orderBy排序策略
	db.Debug().Order("age desc").Find(&studentList)
	fmt.Println(studentList)
	db.Debug().Order("id desc, age desc").Find(&studentList)
	fmt.Println(studentList)

	// limit 每页容量 offset 页面位置     分页
	db.Debug().Limit(2).Offset(0).Find(&studentList)
	fmt.Println(studentList)

	// 先排序再分页

	db.Debug().Order("id desc").Limit(2).Offset(0).Find(&studentList)
	fmt.Println(studentList)

	// 先条件筛选再分页

	db.Debug().Where("name = ?", "yangyang02").Limit(2).Offset(0).Find(&studentList)
	fmt.Println(studentList)

}

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `gorm:"colum:name"`
	Age  int    `gorm:"column:age"`
	int
	dask string
}

func main() {

	student := Student{
		Name: "yang",
		Age:  10,
		int:  1000,
		dask: "111",
	}

	s := reflect.TypeOf(student)

	num := s.NumField()

	for i := 0; i < num; i++ {
		field := s.Field(i)
		fmt.Printf("num: %d, name: %s, offset: %d,anonymous: %t, type: %s, is: %t, gormtag: %s ", i, field.Name,
			field.Offset,       // 偏移量
			field.Anonymous,    // 是否匿名
			field.Type,         // 类型
			field.IsExported(), // 是否可导出
			field.Tag.Get("gorm"))

		fmt.Println()
	}

}

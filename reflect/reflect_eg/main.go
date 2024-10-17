package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name   string `gorm:"colum:name"`
	Age    int    `gorm:"column:age"`
	Number int    `gorm:"column:number"`
	num    int
	MMMAge int
}

func main() {

	fields := GetTableFields(User{})

	fmt.Println(fields)

}

func GetTableFields(user User) []string {

	// gorm:"column:name;size:255;not null"  实现取name
	t := reflect.TypeOf(user)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {

		// 定义结果切片
		res := make([]string, 0, t.NumField())

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.IsExported() {
				name := Camel2Snake(field.Name)
				if len(field.Tag.Get("gorm")) > 0 {

					content := field.Tag.Get("gorm")
					if strings.HasPrefix(content, "column:") {
						content = content[7:]
						index := strings.Index(content, ";")
						if index > 0 {
							name = content[:index]
						} else {
							name = content
						}
					}
				}
				res = append(res, name)
			}
		}
		return res
	} else {
		return nil
	}
}

func IsUpper(c byte) bool {

	return c <= 'A' || c <= 'Z'

}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func Camel2Snake(s string) string {
	resS := make([]byte, 0, len(s)+4)
	if len(s) == 0 {
		return ""
	}

	if IsUpper(s[0]) {
		resS = append(resS, UpperLowerExchange(s[0]))
	} else {
		resS = append(resS, s[0])
	}

	for i := 1; i < len(s); i++ {
		c := s[i]
		if IsUpper(c) {
			resS = append(resS, '_', UpperLowerExchange(c))
		} else {
			resS = append(resS, c)
		}
	}

	return string(resS)
}

func UpperLowerExchange(c byte) byte {
	return c ^ ' '
}

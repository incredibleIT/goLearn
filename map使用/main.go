package main

import (
	"fmt"
)

func main() {

	var student = map[int]map[string]string{1001: {"name": "yang", "age": "18"}, 1002: {"name": "yuan", "age": "19"}}
	fmt.Println(student)

	var data = make(map[string][]string)
	data["shandong"] = []string{"weihai", "jinan", "heze", "yantai"}
	data["hebei"] = []string{"shijiazhuang", "beijing"}
	fmt.Println(data["hebei"][1])
	fmt.Println(data)

	// 切片是map类型
	a := make([]map[string]string, 3)
	a[0] = map[string]string{"name": "yang", "age": "18"}
	a[1] = map[string]string{"name": "yang", "age": "19"}
	a[2] = map[string]string{"name": "da", "age": "19"}
	fmt.Println(a)

	// range切片map
	for _, v := range a {

		fmt.Println(v)

		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}

	}

}

package main

import (
	"fmt"
)

func main() {
	data := map[int]map[string]string{
		1001: {"name": "yyy", "yvwen": "111"},
		1002: {"name": "dada", "yvwen": "222"},
		1003: {"name": "zhu", "yvwen": "333"},
		1004: {"name": "zhu2", "yvwen": "444"},
		1005: {"name": "ggg", "yvwen": "555"},
	}

	for _, v := range data {
		fmt.Println(v)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}

	fmt.Println("---------------")

	delete(data, 1004)

	for _, v := range data {
		fmt.Println(v)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}

	for i, j := range "hello world" {
		fmt.Println(i, string(j))
	}

}

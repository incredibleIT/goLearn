package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	slcB, _ := json.Marshal([]string{"apple", "peach", "pear"})
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]string{"apple": "peach", "pear": "peach"})
	fmt.Println(string(mapB))

	mapC, _ := json.Marshal(map[string]int{"apple": 1, "pear": 2})

	res1, _ := json.Marshal(response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}})
	fmt.Println(string(res1))

	res2, _ := json.Marshal(response2{Page: 1, Fruits: []string{"apple", "peach", "pear"}})
	fmt.Println(string(res2))

	var m map[string]interface{}

	json.Unmarshal(mapC, &m)
	fmt.Println(m)

	fmt.Println(runtime.NumCPU())

}

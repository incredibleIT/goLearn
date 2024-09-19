package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4}

	//追加元素
	arr1 := append(arr, 5, 6)
	fmt.Println(arr1)

	// 追加切片 (多个元素)
	arr2 := append(arr1, []int{7, 8, 9}...)
	fmt.Println(arr2)

	// 删除元素, 例如删除索引为2
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	sliceFirst := slice[:index]
	sliceSecond := slice[index+1:]
	slice = append(sliceFirst, sliceSecond...)
	fmt.Println(slice) // [1, 2, 4, 5]

	// 插入
	// 在开头位置插入元素
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(append([]int{6}, slice1...))
	// 在任意位置插入元素

	// 错误示范
	slice2 := []int{1, 2, 3, 4, 5}
	index2 := 2
	//此时 slice2[:index2]容量足够向其中追加7时会覆盖slice2的后面元素
	//fmt.Println(append(append(slice2[:index2], []int{7}...), slice2[index2:]...))

	// 正确
	fmt.Println(append(slice2[:index2], append([]int{7}, slice2[index2:]...)...))

	// 排序sort
	slice3 := []int{2, 423, 4, 234, 234, 5, 234, 32, 321, 4, 234, 234, 234}
	sort.Ints(slice3)
	fmt.Println(slice3)

}

package main

import "fmt"

func main() {

	// 定义一个长度是5, 容量是10的切片
	s1 := make([]int, 5, 10)
	// 对s1扩容赋值给s2, 此时s2长度6, 容量10
	s2 := append(s1, 555)

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// 若此时更改s1 例如s1[0] = 100, 此时s2会相应改变
	//s1与s2依然存在联系是不希望看到的
	s1[0] = 100
	fmt.Println(s2)

	// 扩容机制
	//a切片长度5, 容量5
	var a = []int{1, 2, 3, 4, 5}
	// 当对a进行扩容时, 容量不足以扩容, 根据扩容机制就会对原切片进行拷贝之后容量变为原切片二倍赋给新切片, 但二倍也不是必然的go也会考虑内存对齐等情况导致偏差
	// 所以此时a2长度为6, 容量为12(考虑了其他情况)
	a2 := append(a, 555)

	fmt.Println(a2, len(a2), cap(a2))
	fmt.Println(len(a), cap(a))

	var b1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b2 := append(b1, 11)
	fmt.Println(b2, len(b2), cap(b2))

}

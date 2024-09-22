package main

import "fmt"

// return 语句不是原子操作, 而是被拆分了两步
// 将返回值赋值rval, 然后ret返回
// 而defer语句在这两条语句之间执行

func f1() int {
	i := 5
	defer func() {
		i++
	}()

	return i
}

func f2() *int {
	i := 5
	fmt.Println(&i)
	defer func() {
		i++
		fmt.Println(&i)

	}()

	return &i
}

// 返回值命名会替换掉rval
func f3() (result int) {

	defer func() {
		result++
	}()

	return 5
}

func f4() (result int) {

	defer func() {
		result++
	}()

	return result
}

func f5() (r int) {
	t := 5
	defer func() {
		t++
	}()
	return t
}

func f6() (r int) {

	defer func(r int) {
		r++
	}(r)
	return 5
}

func f7() (r int) {

	defer func(x int) {
		r = x + 1 // 对外部的引用
	}(r)
	return 5

}

func main() {
	//fmt.Println(f1())
	//fmt.Println(*f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
	//fmt.Println(f5())
	//fmt.Println(f6())
	//fmt.Println(f7())
}

package main

// 返回空值
func fun1() {
	return
}

// 返回单值
func fun2() int {
	return 10
}

// 返回多值
func fun3() (string, int) {
	return "hello", 10
}

//返回值命名
// 此时不会返回零值, 返回的默认为a的值, a在函数中并未赋值此时返回的是a的默认值
func fun4() (a int) {
	return a
}

func main() {

}

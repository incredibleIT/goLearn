package main

import "fmt"

func main() {

	//
	var a int
	var b, c int
	var (
		d int
		e string
		f float64
	)

	fmt.Println(a, b, c, d, e, f)

	var g int = 1

	var h = 1

	var i, j = 1, 1

	var (
		k = 1
		y = 1
		l = 1
	)

	m := 1
	x, z := 1, 1
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, y, z, l, m, x)

}

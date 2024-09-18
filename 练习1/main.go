package main

import (
	"fmt"
)

func main() {

	// 1
	first, second := 0, 1

	for num := 0; num < 9; num++ {
		sum := first + second
		first = second
		second = sum
	}
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
	fmt.Println(second)

	// 2

	sum := 0

	for i := 1; i <= 10; i++ {
		num := 1

		for j := 2; j <= i; j++ {
			num *= j
		}

		sum += num
	}

	fmt.Println(sum)

	// 3

	//result := 10
	//var number int
	//
	//for {
	//	n, _ := fmt.Scanln(&number)
	//	fmt.Println(n)
	//	if number < result {
	//		fmt.Println("太大了")
	//	} else if number > result {
	//		fmt.Println("太小了")
	//	} else {
	//		fmt.Println("猜对了")
	//		break
	//	}
	//}

	// 4
	for i := 0; i < 6; i++ {
		for j := 1; j <= 6-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}

	for i := 5; i >= 0; i-- {
		for j := 6 - i; j > 0; j-- {
			fmt.Print(" ")
		}

		for j := 2*i + 1; j > 0; j-- {
			fmt.Print("*")
		}

		fmt.Println()
	}

}

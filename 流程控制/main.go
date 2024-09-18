package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	}

	if 6%2 == 0 {
		fmt.Println("6 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

}

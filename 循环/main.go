package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	for {

		fmt.Println("hello")

		break

	}

	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	s := make(chan int)

	go func() {
		s <- 10
	}()

	time.Sleep(1 * time.Second)

	fmt.Println(<-s)

}

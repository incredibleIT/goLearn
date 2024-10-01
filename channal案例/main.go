package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func producerInt(c chan<- string) {
	var x = ""
	for {
		reader := bufio.NewReader(os.Stdin)
		x, _ = reader.ReadString('\n')
		c <- x
	}
}

func producerString(c chan<- string) {
	for i := 0; i < 1; i++ {
		time.Sleep(5 * time.Second)
		c <- "hello world"
	}
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go producerInt(c1)

	go producerString(c2)

	for {
		select {
		case v1 := <-c1:
			fmt.Println(v1)
		case v2 := <-c2:
			fmt.Println(v2)
		}
	}

}

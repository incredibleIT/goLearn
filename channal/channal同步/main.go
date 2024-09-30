package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("worker start")
	time.Sleep(time.Second * 1)
	fmt.Println("worker end")
	done <- true
}

func main() {

	done := make(chan bool)

	go worker(done)

	<-done

}

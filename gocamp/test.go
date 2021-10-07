package main

import (
	"fmt"
	"time"
)

func test(ch chan int) {
	fmt.Println("start test")
	a := <-ch
	fmt.Println(a)
}

func main() {
	ch := make(chan int)
	go test(ch)
	time.Sleep(2 * time.Second)
	close(ch)
	time.Sleep(10 * time.Second)
}

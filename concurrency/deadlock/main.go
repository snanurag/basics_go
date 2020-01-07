package main

import (
// "fmt"
)

func main() {
	ch := make(chan int)
	go test()
	go test()
	ch <- 1

}

func test() {
	ch := make(chan int)
	ch <- 1

}

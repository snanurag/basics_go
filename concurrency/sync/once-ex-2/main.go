package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	j    int
)

func main() {
	test()
	test()
}
func test() {
	var i int

	once.Do(func() {
		i = 100
		j = 50
	})

	fmt.Println(i)
	fmt.Println(j)
}

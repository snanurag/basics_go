package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)

	for i := 0; i < 100; i++ {
		go func() {
			once.Do(onceBody) // This will print first value.. then will never print again
			fmt.Println(i)    // This will print rest of the values.
		}()
	}
	fmt.Println(999999)
	<-done
}

package main

import (
	"time"
	"fmt"
)

var limit = make(chan int, 3)
var work = []func(){sleep,sleep,sleep,sleep,sleep,sleep,sleep,sleep,sleep,sleep,sleep,sleep}
 
func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}


func sleep(){
	time.Sleep(1000 *time.Millisecond)
	fmt.Println("sleep executed")
}
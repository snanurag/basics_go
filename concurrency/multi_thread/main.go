package main

/**

 If you are on mac then install htop utility to see the core-wise cpu usage.
 This program aims to visualize that binary generated from go run as multi-processor process.
 From htop it can be seen that the number of cores busy is equal to the threads running here.
*/

import (
	"runtime"
)

func main() {

	// GOMAXPROCS allows the maximum possible cores to be consumed if set to -1. So for -1, 7 cores are busy. CPU usage is 700%
	// If you change GOMAXPROCS to 2 then the number of cores busy are 2 only. CPU usage is 200%.
	runtime.GOMAXPROCS(2)
	ch := make(chan int)
	go alwaysTrue()
	go alwaysTrue()
	go alwaysTrue()
	go alwaysTrue()
	go alwaysTrue()
	go alwaysTrue()
	go alwaysTrue()
	print(<-ch)
}

func alwaysTrue(){
	sum := 1
	for sum < 1000 {
		sum =1
	}
}
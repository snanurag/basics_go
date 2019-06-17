package main
/**
Output is mostly 2000000 but sometiems 1999999 because after last send to channel from change,
main func's print is invoked rather than executing increment func.

*/
import (
	"sync"
)

var i =0;
var wg sync.WaitGroup
func main() {
	
	wg.Add(2)
	ch := make(chan int)
	go change(ch)
	go change(ch)
	go increment(ch)
	wg.Wait()
	println(i)	
}

func change(ch chan int) {
	for j:=0; j<1000000; j++ {
		ch <- 1
	}
	wg.Done()
}

func increment(ch chan int){
	for {
		i += <-ch
	}
}
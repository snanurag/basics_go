package main
/**
Here sum is never 2000000. This proves that the 2 threads using different cores.
Every core loads value in its regsiter and updates it back to the memory. 
This 2-step process introduces local copy of variable and hence the requirement of synchronization.
*/
import (
	"sync"
//	"runtime"
)

var i =0;
var wg sync.WaitGroup
func main() {

	// Using single processor causes tremendous effect on the output.
	// Output is exactly 2000000 if using single processor.
//	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go change()
	go change()
	wg.Wait()
	println(i)	
}

func change() {
	for j:=0; j<1000000; j++ {
		i++
	}
	wg.Done()
}
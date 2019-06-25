package main
import (
//	"runtime"
	// "sync"
)
var a, b int
// var wg sync.WaitGroup

func f() {
	a = 1
	b = 2
	// wg.Done()
}


func main() {
//	runtime.GOMAXPROCS(1)
	for {
		// wg.Add(1)
		go f()
		
		if b ==2 && a == 0 {
			println(b ," " , a)
			println("found condition : write order is not reflected during read.")
			break
		}
		println(b ," " , a)
		// wg.Wait()
		a = 0
		b = 0
		

	}
}
package main
/**
Here printing a variable is never less than the changed var value. This proves that goroutine doesn't
have local copy of variable. This wouldn't have been the nature of output in Java.
*/
import (
	"time"
)

var i =0;
func main() {
	go change()
	go print()
	
	time.Sleep(1000 * time.Millisecond)
	println("hi")
}

func change() {
	for j:=0; j<20; j++ {
		i++
		println("c ", i);
		time.Sleep(5 * time.Millisecond);
	}
}

func print(){
	for j:=0; j<20; j++ {
		
		println("p ", i);
		time.Sleep(5 * time.Millisecond);
	}
}
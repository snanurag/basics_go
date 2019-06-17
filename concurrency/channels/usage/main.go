package main

func main(){
	ch := make(chan int)
	go func() {
		println(<-ch)
	}()
	
	ch <- 1
}

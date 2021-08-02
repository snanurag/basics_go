package main

/*
	Try escape analysis using go build -gcflags '-m' ./main.go
	Output :
		./main.go:15:7: can inline test1
		./main.go:19:6: can inline test2
		./main.go:24:6: can inline main
		./main.go:25:6: inlining call to test1
		./main.go:26:7: inlining call to test2
*/

var a = 45
var c = make(chan int)

var b = &a

func test1() {
	a = 30
}

func test2() {
	a = 50
	c <- 0
}

func main() {
	test1()
	go test2()
	b := <-c
	println(b)
	println(a)
}

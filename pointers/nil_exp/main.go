package main

import "fmt"

type tStruct struct {
}

func main() {
	//test(nil) // Not allowed
	var s *tStruct
	test(*s) //Will give error
	var t tStruct
	test(t)

}

func test(t tStruct) {
	fmt.Println("value ", t)
}

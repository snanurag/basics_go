package main

import "fmt"

type add func(a int, b int) int

func (a add) getResult() int {
	return a(5, 6)
}

func getResult2(a add) int {
	return a(3, 4)
}
func main() {

	x := func(a, b int) int {
		return a + b
	}
	var y add = x
	fmt.Println(y.getResult())

	fmt.Println(getResult2(func(a int, b int) int {
		return a + b
	}))
}

package main

import "fmt"

func withPointer() {
	var a [3]int
	b := &a
	b[2] = 7
	fmt.Println(a, *b) // prints [0 0 0 0 0] [0 0 7 0 0]
}
func withoutPointer() {
	var a [5]int
	b := a
	b[2] = 7
	fmt.Println(a, b) // prints [0 0 0 0 0] [0 0 7 0 0]
}

func slice() {
	a := []int{1, 2}
	b := a
	b[0] = 0
	fmt.Println(a, b)
}

func appendToArray() {
	// Can't use append to array
	//a := [3]int{1,2,3}
	//a = append(a, 1)
}

func appendToSlice() {
	a := []int{1, 2, 3}
	a = append(a, 1)
	fmt.Println(a)

}
func main() {
	withoutPointer()
	withPointer()
	slice()
	appendToArray()
	appendToSlice()
}

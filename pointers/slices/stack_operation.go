package main

import "fmt"

func appendVal(s []int) {
	fmt.Println(s)
	s = append(s, 10)
	fmt.Println(s)
}

func stack_op_from_fixed_underlying_array() {
	s := []int{1, 2, 3}
	appendVal(s)
	fmt.Println(s, len(s)) // prints [1 2 3] 3
}

func stack_op_from_slice() {
	var s = []int{1, 2, 3, 4, 5}
	s = s[1:3]
	appendVal(s)
	fmt.Println(s)
	s = s[:4]
	fmt.Println(s)
}

func main() {
	stack_op_from_fixed_underlying_array()
	fmt.Println("-----")
	stack_op_from_slice()
}

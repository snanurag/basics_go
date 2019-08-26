package main

import "fmt"

func negate(s []int) {
	for i := range s {
		s[i] = -s[i]
	}
}

func negation() {
	var a = []int{1, 2, 3, 4, 5}
	negate(a)
	fmt.Println(a) // prints [-1 -2 -3 -4 -5]
}

func main() {
	negation()
}

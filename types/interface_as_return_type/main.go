package main

import "fmt"

type A struct {
	a int
}

func getInt() interface{} {
	return 3
}

func getA() interface{} {
	return &A{30}
}

func main() {
	var b *A
	b = getA().(*A)
	fmt.Println(b)

	fmt.Println(getInt())
	var a A
	a = getInt().(A) // Panic at type casting

	fmt.Println(a)
}

package main

import "fmt"

func main() {

	s1 := string("hey")
	s2 := "hey"

	fmt.Println(s1)
	fmt.Println(s1 == s2)
	fmt.Println(&s1)
	fmt.Println(&s2)

	a := []rune(s1)
	b := []int32(s1)
	fmt.Println(string(a))
	fmt.Println(string(b))

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a[0] == b[0])
}

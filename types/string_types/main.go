package main

import "fmt"

func main() {

	s1 := string("hey")
	s2 := "hey"

	fmt.Println(s1)
	fmt.Println(s1 == s2)
	fmt.Println(&s1)
	fmt.Println(&s2)

}

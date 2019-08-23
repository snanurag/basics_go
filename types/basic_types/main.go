package main

import "fmt"

type MyString string

func main() {

	s := MyString("hi world!")
	v := s
	fmt.Println(s)
	fmt.Println(v)
	fmt.Println(&s)
	fmt.Println(&v)

}

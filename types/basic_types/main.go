package main

import "fmt"

type MyString string
type MyInt int8

func main() {

	s := MyString("hi world!")
	v := s
	fmt.Println(s)
	fmt.Println(v)
	fmt.Println(&s)
	fmt.Println(&v)

	i := MyInt(1)
	j := i
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(&i)
	fmt.Println(&j)

}

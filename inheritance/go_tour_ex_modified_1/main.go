package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
	getStr() string
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	fmt.Println(a.getStr())

	s := MyString("hi")
	a = s
	fmt.Println(a.Abs())
	fmt.Println(a.getStr())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) getStr() string {
	return "Hello world!"
}

type MyString string

func (f MyString) Abs() float64 {
	return 0
}

func (f MyString) getStr() string {
	return "hi world!"
}

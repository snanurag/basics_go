package main

import (
	"fmt"
	"reflect"
)

type Int1 interface {
}

type Strct1 struct {
	s string
}

func main() {
	var i Int1
	i = Strct1{
		"hi",
	}
	fmt.Println(i)
	fmt.Println(reflect.ValueOf(i))

}

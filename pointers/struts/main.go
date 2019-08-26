package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func withoutPointers() {
	v1 := Vertex{1, 2}
	v2 := v1
	v2.Lat = 0
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println("------")
}

func withPointers() {
	v1 := Vertex{1, 2}
	v2 := &v1
	v2.Lat = 0
	fmt.Println(v1)
	fmt.Println(*v2)
}

func main() {
	withoutPointers()
	withPointers()
}

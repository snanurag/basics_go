package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func changeValWithoutPointers() {
	v1 := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	v2 := v1
	v2["Google"] = Vertex{1, 2}
	fmt.Println(v1["Google"])
	fmt.Println(v2["Google"])
	fmt.Println("------")
}

//Doesn't compile

func changeValWithPointers() {
	v1 := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	v2 := &v1
	(*v2)["Google"] = Vertex{1, 2}
	fmt.Println(v1["Google"])
	fmt.Println((*v2)["Google"])
	fmt.Println("------")

}

func accessMapUsingPointers() {
	m := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	v1 := &m
	fmt.Println(v1)
	fmt.Println(*v1)
	fmt.Println((*v1)["Google"])

}
func main() {
	changeValWithoutPointers()
	changeValWithPointers()
	accessMapUsingPointers()
}

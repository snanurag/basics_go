package main

import "fmt"

type Parent interface {
}

type Child struct {
}

func main() {

	// IS-A relationship between struct and interface.
	var p Parent
	p = Child{}
	checkingInterfaceType(p)
	acceptsParent(p)
	c := Child{}
	checkingInterfaceType(c)
	acceptsParent(c)
	fmt.Println("######################")

	// IS-A relationship fails between two structs.
	s2 := Struct2{}
	s2.methodOfStruct1() // Composition
	//acceptsStruct1(s2) // Inheritance failed
	acceptsStruct1(Struct1{}) // This works
}

func checkingInterfaceType(p Parent) {
	switch p.(type) {
	case string:
		fmt.Print("Type is string")
	case Child:
		fmt.Println("Type is Child.")
	case Parent:
		fmt.Println("Type is Parent")

	}

	switch p.(type) {
	case string:
		fmt.Print("Type is string")
	case Parent:
		fmt.Println("Type is Parent")
	case Child:
		fmt.Println("Type is Child.")

	}
}

func acceptsParent(p Parent) {

}

type Struct1 struct {
}

func (s *Struct1) methodOfStruct1() {
	fmt.Println("Method from Struct1.")
}

type Struct2 struct {
	Struct1
}

func acceptsStruct1(s Struct1) {

}

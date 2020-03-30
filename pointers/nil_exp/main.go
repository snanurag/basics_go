package main

import "fmt"

type tStruct struct {
}

func main() {
	passPointer(nil) // Can pass nil. That's why not safe.

	var s *tStruct
	fmt.Println("default of pointer is ", s)
	//passValue(*s) //Will give nil dereference error. You don't know if pointer has value or nil

	//passValue(nil) //Compilation error. Can't pass nil in value.
	var t tStruct
	passValue(t)

	var v1 [2]tStruct
	fmt.Println("default of array is empty array", v1)

	var v2 [2]*tStruct
	fmt.Println("default of array of pointer is empty array with nil values", v2)

	var v3 []tStruct
	fmt.Println("default of slice is slice with 0 value ", v3)
	fmt.Println("Operations  allowed on slice", len(v3))

	v3 = make([]tStruct, 2)
	fmt.Println("Initialized slice with default values ", v3)

	var v4 []*tStruct
	fmt.Println("default of slice with pointer values ", v4)
	v4 = make([]*tStruct, 2)
	fmt.Println("Initialized slice with nil values ", v4)
}

func passValue(t tStruct) {
	fmt.Println("value ", t) // Default value is empty struct. Not nil
}

func passPointer(t *tStruct) {

}

func returnNilSlice() []tStruct {
	// Allowed
	return nil
}

func returnNilArray() [2]tStruct {
	//return nil // Not Allowed arrays default type is empty array
	var v [2]tStruct
	return v
}

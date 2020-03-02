package main

import "fmt"

func main() {
	s1 := sampleStruct{
		1,
	}

	fmt.Println(&s1)
	fmt.Println(&s1.i)
	fmt.Println("####################")
	printPointers(s1)
	fmt.Println("####################")
	printValue(&s1)

	fmt.Println("###### With Pointers as var #######")

	i := 1
	s2 := sampleStructWithPt{
		&i,
	}

	fmt.Println(&s2)
	fmt.Println(s2.i)
	fmt.Println("####################")
	printPointers2(s2)
	fmt.Println("####################")
	printValue2(&s2)
}

type sampleStruct struct {
	i int
}

type sampleStructWithPt struct {
	i *int
}

func printPointers(a sampleStruct) {
	fmt.Println(&a)
	fmt.Println(&a.i)
}

func printValue(a *sampleStruct) {
	fmt.Println(a)
	fmt.Println(&a.i)
}

func printPointers2(a sampleStructWithPt) {
	fmt.Println(&a)
	fmt.Println(a.i)
}

func printValue2(a *sampleStructWithPt) {
	fmt.Println(a)
	fmt.Println(a.i)
}

package main

import "fmt"

type hello struct {
	a string
	b int32
}

func main() {
	h1 := hello{
		a: "hi",
		b: 0,
	}

	h2 := hello{
		a: "hi",
		b: 0,
	}

	fmt.Println(&h1 == &h2) // false
	fmt.Println(h1 == h2)   //true
	fmt.Printf("h1 and h2 pointer : %s %s", &h1, &h2)
	fmt.Println()

	h2.b = 1
	fmt.Println(h1 == h2)

}

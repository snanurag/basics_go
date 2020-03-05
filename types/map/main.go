package main

import "fmt"

func main() {
	m := make(map[int]int)

	//Double value return
	if i, ok := m[3]; !ok {
		fmt.Println("not present", i)
	}

	// Single value return
	i := m[3]
	fmt.Println(i)

	//j := t() Single value return does not work for functions.

}

func t() (int, int) {
	return 0, 0
}

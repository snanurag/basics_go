package main

import "fmt"

func main() {
	var m map[string]string
	//m["hi"] = "hi" //By default maps are nil
	fmt.Println(m)

	m = make(map[string]string)
	m["hello"] = "hello"
	fmt.Println(m)

	m = nil
	fmt.Println(m)
	fmt.Println(m["hi"])

}

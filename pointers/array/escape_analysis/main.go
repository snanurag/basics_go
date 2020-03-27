package main

import "fmt"

/**
Do escape analysis using the following command.
go build -gcflags '-m' ./main.go
*/
func main() {

	var u [8192]int
	u = [8192]int{}
	u[0] = 0
	go values(u)
	fmt.Print(u[0])
}

func values(t [8192]int) {
	t[0] = 1
	fmt.Println(t[0])
}

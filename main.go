package main

func main() {

	u := make([]int, 8191) // Does not escape to heap
	_ = u

	v := make([]int, 1) // Escapes to heap == 8 bytes * 2^ 13 = 2 ^16 = 64kb
	_ = v

	var w [1024 * 1024 * 1.25]int
	_ = w
}

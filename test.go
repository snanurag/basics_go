package main

func main() {
	u := make([]int, 8191) // Does not escape to heap
	u = append(u, 0)

	v := make([]int, 8192) // Escapes to heap == 8 bytes * 2^ 13 = 2 ^16 = 64kb
	v = append(v, 0)

	w := make([]int32, 16383) // Does not escape to heap
	w = append(w, 0)

	x := make([]int32, 16384) // Escapes to heap = 2^16 = 64 kb
	x = append(x, 0)

}

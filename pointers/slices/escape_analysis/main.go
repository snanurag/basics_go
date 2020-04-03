package main

/**
Do escape analysis using the following command.
go build -gcflags '-m' ./fib_test.go
*/
func main() {
	//var t []int
	t := []int{}
	t = append(t, 0)
	t = append(t, 0)

	u := make([]int, 8191) // Does not escape to heap
	u = append(u, 0)

	v := make([]int, 8192) // Escapes to heap == 8 bytes * 2^ 13 = 2 ^16 = 64kb
	v = append(v, 0)

	w := make([]int32, 16383) // Does not escape to heap
	w = append(w, 0)

	x := make([]int32, 16384) // Escapes to heap = 2^16 = 64 kb
	x = append(x, 0)

	y := make([]*int, 8191)
	i := 0
	y = append(y, &i)

	z := make([]*int, 8192)
	z = append(z, &i)

	u2 := make([]int, 1) // Escapes to heap because passing into goroutine.
	u2[0] = 1
	u2 = append(u2, 0)
	go values(u2)
	u3 := make([]string, 4000)
	u3 = append(u3, "")
}

func values(t []int) {
}

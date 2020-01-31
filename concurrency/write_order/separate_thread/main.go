package main

var a, b int

func f() {
	a = 1
	b = 2
}

func main() {
	//runtime.GOMAXPROCS(1)
	for {
		go f()

		if b == 2 && a == 0 {
			println(b, " ", a)
			println("found condition : write order is not reflected during read.")
			break
		}
		println(b, " ", a)
		a = 0
		b = 0

	}
}

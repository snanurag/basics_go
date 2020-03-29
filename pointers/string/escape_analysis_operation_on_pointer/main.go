package main

func main() {
	s := "hello world"
	values(&s)
	t := "hi world"
	values2(&t)
}

func values(a *string) {
	x := *a
	_ = x
}

func values2(a *string) {
	var y []*string
	y = append(y, a)
}

package main

/**
Complete escape analysis
go build -gcflags '-m -l' ./fib_test.go
*/

func main() {
	s := "hello world"
	values(&s)
}

func values(s *string) {
	values3(s)
	//go values3(s)
	values4(*s)
}

func values3(s *string) *string {
	return s
}

func values4(s string) *string {
	//fmt.Println(&s)
	return &s
}

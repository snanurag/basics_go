package main

/**
go build -gcflags '-m -l' ./fib_test.go
*/
func main() {
	s := "hello world" // Moved to heap
	go values(&s)

	s2 := "hi world" // Does not move to heap
	go values2(s2)
}

func values(s *string) {
}

func values2(s string) {

}

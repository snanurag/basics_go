package main

/**
 goroutine % go build -gcflags '-m' ./withoutmain.go
# command-line-arguments
./withoutmain.go:28:6: can inline newGoRoutineWithPointer1
./withoutmain.go:31:6: can inline newGoRoutineWithValue1
./withoutmain.go:35:6: can inline newGoRoutineWithIntPointer1
./withoutmain.go:28:31: t does not escape
./withoutmain.go:35:34: t does not escape
./withoutmain.go:20:2: moved to heap: i
./withoutmain.go:18:8: &TestStruct1 literal escapes to heap
*/
type TestStruct1 struct {
}

func test() {
	t1 := &TestStruct1{}
	t2 := TestStruct1{}
	i := 1
	t3 := &i
	go newGoRoutineWithPointer1(t1)
	go newGoRoutineWithValue1(t2)
	go newGoRoutineWithIntPointer1(t3)
}

func newGoRoutineWithPointer1(t *TestStruct1) {
}

func newGoRoutineWithValue1(t TestStruct1) {

}

func newGoRoutineWithIntPointer1(t *int) {

}

package main

/*
 go build -gcflags '-m' ./main.go
# command-line-arguments
./main.go:27:6: can inline newGoRoutineWithPointer
./main.go:30:6: can inline newGoRoutineWithValue
./main.go:34:6: can inline newGoRoutineWithIntPointer
./main.go:27:30: t does not escape
./main.go:34:33: t does not escape
./main.go:20:2: moved to heap: i
./main.go:18:8: &TestStruct literal escapes to heap
*/
type TestStruct struct {
}

func main() {
	t1 := &TestStruct{}
	t2 := TestStruct{}
	i := 1
	t3 := &i
	go newGoRoutineWithPointer(t1)
	go newGoRoutineWithValue(t2)
	go newGoRoutineWithIntPointer(t3)
}

func newGoRoutineWithPointer(t *TestStruct) {
}

func newGoRoutineWithValue(t TestStruct) {

}

func newGoRoutineWithIntPointer(t *int) {

}

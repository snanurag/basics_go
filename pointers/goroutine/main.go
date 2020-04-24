package main

type TestStruct struct {
}

func main() {
	t1 := &TestStruct{}
	t2 := TestStruct{}
	go newGoRoutineWithPointer(t1)
	go newGoRoutineWithValue(t2)
}

func newGoRoutineWithPointer(t *TestStruct) {
}

func newGoRoutineWithValue(t TestStruct) {

}

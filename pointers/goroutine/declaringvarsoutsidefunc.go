package main

/**
If variable is declared outside the function, then it doesn't escape to heap on goroutine start.
go build -gcflags '-m' ./declaringvarsoutsidefunc.go
# command-line-arguments
./declaringvarsoutsidefunc.go:28:6: can inline newGoRoutineWithPointer2
./declaringvarsoutsidefunc.go:31:6: can inline newGoRoutineWithValue2
./declaringvarsoutsidefunc.go:35:6: can inline newGoRoutineWithIntPointer2
./declaringvarsoutsidefunc.go:28:31: t does not escape
./declaringvarsoutsidefunc.go:35:34: t does not escape
*/

var t1 = &TestStruct2{}
var t2 = TestStruct2{}
var i = 1
var t3 = &i

type TestStruct2 struct {
}

func test2() {
	go newGoRoutineWithPointer2(t1)
	go newGoRoutineWithValue2(t2)
	go newGoRoutineWithIntPointer2(t3)
}

func newGoRoutineWithPointer2(t *TestStruct2) {
}

func newGoRoutineWithValue2(t TestStruct2) {

}

func newGoRoutineWithIntPointer2(t *int) {

}

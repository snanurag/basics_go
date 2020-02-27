package main

type Int1 interface {
	IntFunc()
	printImpl()
}

type Impl struct {
	instance1 string
	instance2 string
}

func (t *Impl) IntFunc() {
	println("Impl IntFunc")
}

//func (t *Impl) printImpl() {
//	print("Impl print")
//}

func main() {

	t := Impl{}
	t.IntFunc()

	// This is not working because *Impl does not implement all of Int1 functions.
	//var u Int1
	//u = &Impl{}

	// This is also not working even after implementing the unimplemented functions because
	// functions are implemented on the pointer no actual struct.
	//var v Int1
	//v = Impl{}

}

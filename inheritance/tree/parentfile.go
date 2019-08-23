package tree

type Int1 interface {
	IntFunc()
	printImpl()
}

type Impl struct {
	instance1 string
	instance2 string
}

func printsomething() {

}

func (t *Impl) IntFunc() {
	println("Impl IntFunc")
}

//func (t *Impl) printImpl() {
//	print("Impl print")
//}

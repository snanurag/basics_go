package tree

type Int1 interface {
	IntFunc()
}

type Impl struct {
	instance1 string
	instance2 string
}

func printsomething() {

}

func (t *Impl) IntFunc() {
	print("Impl IntFunc")
}

func (t *Impl) printImpl() {
	print("Impl print")
}

package src

import "fmt"

type parent struct {
}

type Impl struct {
	parent
}

func (p *parent) PrintSomethingFromParent() {
	fmt.Println("Hi from parent.")
}

func (t *Impl) IntFunc() {
	println("Impl IntFunc")
}

//func (t *Impl) printImpl() {
//	print("Impl print")
//}

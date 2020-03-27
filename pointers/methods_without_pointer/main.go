package main

import "fmt"

type Person struct {
	name string
}

func main() {
	var p Person
	p.changeName()
	p.printName()
	p.changeName2()
	p.printName2()
}

func (p Person) changeName() {
	p.name = "NewName"
}

func (p Person) printName() {
	fmt.Println("Name from value method : ", p.name)
}

func (p *Person) changeName2() {
	p.name = "NewName"
}

func (p *Person) printName2() {
	fmt.Println("Name from value method : ", p.name)
}

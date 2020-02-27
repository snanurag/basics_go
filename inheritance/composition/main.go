package main

import "fmt"

type Driver struct {
	*Controller
}

type Controller struct {
}

func (d *Controller) hey() {
	fmt.Print("Hey Controller!!")
}

func main() {
	d := &Driver{}
	d.hey()
}

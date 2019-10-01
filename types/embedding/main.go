package main

import "fmt"

type Driver struct {
	*Controller
}

type Controller struct {
}

func (d *Driver) hey() {
	fmt.Print("Hey Driver!!")
}

func main() {
	d := &Driver{}
	d.hey()
}

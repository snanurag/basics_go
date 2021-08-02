package main

import "errors"

func main() {
	err := test()
	print(err.Error())

}

func test() error {
	return errors.New("test err")
	//return nil     // Replacing it will return nil will give panic.
}

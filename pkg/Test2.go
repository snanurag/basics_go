package pkg

import "fmt"

type Test struct {
	hi string
}

func (t *Test) New() error {
	fmt.Print("From Test2")

	return nil
}

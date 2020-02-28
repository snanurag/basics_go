package main

import (
	"fmt"
	"github.com/snanurag/basics_go/collections"
)

func main() {
	maxHeapExample()
}

func maxHeapExample() {
	h := collections.NewHeap()

	h.Push(3)
	h.Push(2)
	h.Push(7)

	for h.Len() > 0 {
		fmt.Println(h.Peek())
		fmt.Println(h.Pop())
	}
}

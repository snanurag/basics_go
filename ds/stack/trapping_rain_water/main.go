package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	fmt.Println(trap([]int{2, 0, 1, 2}) == 3)
	min(1, 2)
}
func trap(height []int) int {
	var ans, current int
	st := stack.New()
	for current < len(height) {
		for st.Len() != 0 && height[current] > height[st.Peek().(int)] {
			top := st.Peek().(int)
			st.Pop()
			if st.Len() == 0 {
				break
			}
			distance := current - st.Peek().(int) - 1
			bounded_height := min(height[current], height[st.Peek().(int)]).(int) - height[top]
			ans += distance * bounded_height
		}
		st.Push(current)
		current++
	}
	return ans
}

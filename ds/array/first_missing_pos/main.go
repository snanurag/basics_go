package main

import "fmt"

func firstMissingPositive(nums []int) int {
	m := make(map[int]int)
	leastPos := 1

	for _, e := range nums {
		m[e] = 1
		for _, ok := m[leastPos]; ok; {
			leastPos++
			_, ok = m[leastPos]
		}
	}

	return leastPos
}

func main() {
	fmt.Println(firstMissingPositive([]int{2, 3, -3, 4, 1}))
}

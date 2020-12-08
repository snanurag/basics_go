// Problem https://www.interviewbit.com/problems/max-distance/

package main

import "fmt"

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func maximumGap(A []int) int {
	if len(A) < 1 {
		return -1
	}
	if len(A) == 1 {
		return 0
	}

	lmin := make([]int, 0)
	rmax := make([]int, len(A))
	min, max := A[0], A[len(A)-1]
	sol := -1

	for i := 0; i < len(A); i++ {
		if A[i] <= min {
			min = A[i]
		}
		lmin = append(lmin, min)
	}

	for i := len(A) - 1; i >= 0; i-- {
		if A[i] >= max {
			max = A[i]
		}
		rmax[i] = max
	}

	fmt.Println(lmin)
	fmt.Println(rmax)
	i, j := 0, 0
	for j < len(A) && i < len(A) {
		if lmin[i] <= rmax[j] {
			sol = mx(sol, j-i)
			j = j + 1
			fmt.Printf("One sol %d %d", i, j)
			fmt.Println()
		} else {
			i = i + 1
			fmt.Printf("One inc %d %d", i, j)
			fmt.Println()
		}
	}
	return sol
}

func mx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumGap([]int{14, 10, 2, 13, 1}))
}

// https://www.interviewbit.com/problems/balance-array/
package main

import "fmt"

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func solve(A []int) int {
	leftEven := make([]int, len(A))
	leftOdd := make([]int, len(A))
	rightEven := make([]int, len(A))
	rightOdd := make([]int, len(A))

	if len(A) == 1 {
		return 1
	}

	for i := 0; i < len(A); i++ {
		if i == 0 {
			leftEven[i] = A[i]
		} else if i%2 == 1 {
			leftOdd[i] = A[i] + leftOdd[i-1]
			leftEven[i] = leftEven[i-1]
		} else {
			leftEven[i] = leftEven[i-1] + A[i]
			leftOdd[i] = leftOdd[i-1]
		}
	}

	for i := len(A) - 1; i >= 0; i-- {
		if i == len(A)-1 {
			if i%2 == 0 {
				rightEven[i] = A[i]
			} else {
				rightOdd[i] = A[i]
			}
			continue
		}
		if i%2 == 0 {
			rightEven[i] = A[i] + rightEven[i+1]
			rightOdd[i] = rightOdd[i+1]
		} else if i%2 == 1 {
			rightOdd[i] = A[i] + rightOdd[i+1]
			rightEven[i] = rightEven[i+1]
		}
	}

	var count int
	for i := 1; i < len(A)-1; i++ {
		if leftEven[i-1]+rightOdd[i+1] == leftOdd[i-1]+rightEven[i+1] {
			count++
		}
	}
	if rightOdd[1] == rightEven[1] {
		count++
	}
	if leftEven[len(A)-2] == leftOdd[len(A)-2] {
		count++
	}
	return count
}

func main() {
	fmt.Println(solve([]int{2, 8, 2, 2, 6, 3}))
}

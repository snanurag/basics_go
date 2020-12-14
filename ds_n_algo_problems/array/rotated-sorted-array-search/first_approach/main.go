// https://www.interviewbit.com/problems/rotated-sorted-array-search/

package main

import (
	"fmt"
)

func main() {
	v := search([]int{1, 7, 67, 133, 178}, 1)
	fmt.Println(v)
}

func search(A []int, B int) int {
	var l, r int
	l = 0
	r = len(A) - 1
	for true {
		i := (l + r) / 2
		if A[i] >= A[0] {
			l = i
		} else {
			r = i
		}
		if r-l == 1 {
			break
		}
	}

	if A[l] > A[r] {
		al := A[0:r]
		ar := A[r:]
		s, v := BinarySearch(B, al)
		if s {
			return v
		}
		s, v = BinarySearch(B, ar)
		if s {
			return len(al) + v
		} else {
			return -1
		}
	} else {
		s, v := BinarySearch(B, A)
		if s {
			return v
		} else {
			return -1
		}
	}
}

func BinarySearch(needle int, haystack []int) (bool, int) {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false, -low
	}

	return true, low
}

// https://www.interviewbit.com/problems/rotated-sorted-array-search/
package main

import "fmt"

func search(A []int, B int) int {

	if A[0] < A[len(A)-1] {
		s, v := BinarySearch(B, A)
		if s {
			return v
		} else {
			return -1
		}
	}

	var l, r int
	l = 0
	r = len(A) - 1
	for A[l] > A[r] {
		mid := (l + r) / 2
		if A[mid] > A[0] {
			if B <= A[mid] && B >= A[l] {
				s, v := BinarySearch(B, A[l:mid+1])
				if s {
					return l + v
				} else {
					return -1
				}
			} else {
				l = mid
			}
		} else {
			if B >= A[mid] && B <= A[r] {
				s, v := BinarySearch(B, A[mid:r+1])
				if s {
					return mid + v
				} else {
					return -1
				}
			} else {
				r = mid
			}
		}
	}
	return -1
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

func main() {
	v := search([]int{1, 7, 67, 133, 178}, 1)
	fmt.Println(v)
}

package main

import "fmt"

/**
https://leetcode.com/problems/median-of-two-sorted-arrays/
*/

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{}))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length == 0 {
		return 0
	}

	var i int
	if length%2 == 0 {
		i = length / 2
	} else {
		i = length/2 + 1
	}

	var j, k, l, n int = 0, 0, 0, 0

	for j < i {
		n = getNextNum(&k, &l, &j, nums1, nums2)
	}

	if length%2 == 0 {
		n += getNextNum(&k, &l, &j, nums1, nums2)
		return float64(n) / 2
	} else {
		return float64(n)
	}

}

func getNextNum(k *int, l *int, j *int, nums1 []int, nums2 []int) int {
	var n int
	if *k < len(nums1) && *l < len(nums2) {
		if nums1[*k] < nums2[*l] {
			n = nums1[*k]
			*k++
		} else {
			n = nums2[*l]
			*l++
		}
	} else if *l == len(nums2) {
		n = nums1[*k]
		*k++
	} else {
		n = nums2[*l]
		*l++
	}
	*j++
	return n
}

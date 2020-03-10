package main

/**
https://leetcode.com/problems/sliding-window-maximum/
*/
import (
	"fmt"
	treeset "github.com/emirpasic/gods/sets/treeset"
)

func main() {
	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {

	m := make(map[int]int)
	var r []int
	s := treeset.NewWithIntComparator()
	i := 0
	for s.Size() < k && i < len(nums) {
		s.Add(nums[i])
		m[nums[i]]++
		//fmt.Println(s)
		if i >= k-1 {
			val, ok := getMaxValue(s)
			if ok {
				r = append(r, val)
			} else {
				return nil
			}
			if m[nums[i-k+1]]--; m[nums[i-k+1]] == 0 {
				s.Remove(nums[i-k+1])
			}
		}
		i++
	}
	return r
}

func getMaxValue(s *treeset.Set) (int, bool) {
	it := s.Iterator()
	it.End()
	if it.Prev() {
		return it.Value().(int), true
	} else {
		return -1, false
	}

}

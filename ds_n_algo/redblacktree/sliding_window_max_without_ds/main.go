package main

/**
https://leetcode.com/problems/sliding-window-maximum/
*/
import (
	"fmt"
	"github.com/snanurag/basics_go/collections/rbt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
func maxSlidingWindow(nums []int, k int) []int {

	var r []int
	m := make(map[int]int)
	s := rbt.NewWithIntComparator()
	i := 0
	for s.Size() < k && i < len(nums) {
		s.Put(nums[i], nil)
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

func getMaxValue(s *rbt.Tree) (int, bool) {
	it := s.Iterator()
	it.End()
	if it.Prev() {
		return it.Key().(int), true
	} else {
		return -1, false
	}

}

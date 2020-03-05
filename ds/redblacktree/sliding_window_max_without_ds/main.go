package main

import (
	"fmt"
	"github.com/snanurag/basics_go/collections/rbt"
	"math/rand"
)

func main() {
	maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
}
func maxSlidingWindow(nums []int, k int) []int {

	var r []int
	s := rbt.NewWithIntComparator()
	i := 0
	for s.Size() < k && i < len(nums) {
		s.Put(nums[i], rand.Int())
		fmt.Println(s)
		if s.Size() == k {
			val, ok := getMaxValue(s)
			if ok {
				r = append(r, val)
			} else {
				return nil
			}
			s.Remove(nums[i-k+1])
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

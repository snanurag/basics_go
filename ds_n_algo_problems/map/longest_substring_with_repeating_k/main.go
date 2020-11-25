// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
package main

import "fmt"

func longestSubstring(s string, k int) int {
	m := make(map[string]int)
	for _, v := range s {

		m[string(v)] += 1
		fmt.Println(m[string(v)])
	}
	return 0
}

func main() {
	longestSubstring("hhhi", 1)
}

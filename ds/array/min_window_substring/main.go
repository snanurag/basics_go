package main

import "fmt"

/**
https://leetcode.com/problems/minimum-window-substring/
*/

func main() {
	m := make(map[string]int)
	fmt.Println(minWindow("a", "a"))
	fmt.Println(m["t"])
}

func minWindow(s string, t string) string {
	m := make(map[string]int)
	orig := make(map[string]int)
	for _, r := range t {
		orig[string(r)] = 1
	}
	for _, r := range s {
		m[string(r)] = -1
	}
	var p1, p2, s1, s2, count int
	for p2 < len(s) {
		if charExistsInSubStr(s[p2], orig) {
			if m[string(s[p2])] == -1 {
				count++
			}
			m[string(s[p2])]++
		}
		if count != len(t) {
			p2++
			continue
		} else {
			if s2 == 0 || p2-p1 < s2-s1 {
				s1 = p1
				s2 = p2
			}
		}
		for p1 < len(s) {
			if m[string(s[p1])] > -1 {
				m[string(s[p1])]--
			}
			if charExistsInSubStr(s[p1], orig) && m[string(s[p1])] == -1 {
				count--
				break
			}
			if p2-p1 < s2-s1 {
				s1 = p1
				s2 = p2
			}
			p1++
		}
		if p2-p1 < s2-s1 {
			s1 = p1
			s2 = p2
		}
		p1++
		p2++
	}
	return s[s1 : s2+1]

}

func charExistsInSubStr(c uint8, m map[string]int) bool {
	if m[string(c)] == 1 {
		return true
	} else {
		return false
	}
}

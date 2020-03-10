package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(anagrams([]string{"cat", "dog", "god", "tca"}))
}
func anagrams(A []string) [][]int {
	m := make(map[string][]int)
	var f [][]int
	for i, e := range A {
		e = getKeyStr(e)
		if m[e] == nil {
			m[e] = []int{i + 1}
		} else {
			m[e] = append(m[e], i+1)
		}

	}
	for _, v := range m {
		f = append(f, v)
	}

	return f
}

func getKeyStr(a string) string {
	a = strings.ToLower(a)
	s := strings.Split(a, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

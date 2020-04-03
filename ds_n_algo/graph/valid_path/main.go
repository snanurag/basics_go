package main

/**
https://www.interviewbit.com/problems/valid-path/
*/

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print(solve(2, 3, 1, 1, []int{2}, []int{3}))

}
func solve(x int, y int, N int, r int, E []int, F []int) string {
	visited := make([][]bool, x+1)
	for i, _ := range visited {
		visited[i] = make([]bool, y+1)
	}

	if traversal(0, 0, x, y, r, E, F, visited) {
		return "YES"
	} else {
		return "NO"
	}

}

func traversal(a, b, x, y int, r int, e, f []int, visited [][]bool) bool {

	if a == x && b == y {
		return true
	}
	if visited[a][b] {
		return false
	} else {
		visited[a][b] = true
	}
	if feasible(a+1, b+1, x, y, r, e, f) && traversal(a+1, b+1, x, y, r, e, f, visited) {
		return true
	} else if feasible(a-1, b+1, x, y, r, e, f) && traversal(a-1, b+1, x, y, r, e, f, visited) {
		return true
	} else if feasible(a+1, b-1, x, y, r, e, f) && traversal(a+1, b-1, x, y, r, e, f, visited) {
		return true
	} else if feasible(a-1, b-1, x, y, r, e, f) && traversal(a-1, b-1, x, y, r, e, f, visited) {
		return true
	} else if feasible(a, b+1, x, y, r, e, f) && traversal(a, b+1, x, y, r, e, f, visited) {
		return true
	} else if feasible(a, b-1, x, y, r, e, f) && traversal(a, b-1, x, y, r, e, f, visited) {
		return true
	} else if feasible(a+1, b, x, y, r, e, f) && traversal(a+1, b, x, y, r, e, f, visited) {
		return true
	} else if feasible(a-1, b, x, y, r, e, f) && traversal(a-1, b, x, y, r, e, f, visited) {
		return true
	}
	return false
}

func feasible(a, b, x, y, r int, e, f []int) bool {
	if inRect(a, b, x, y) && !liesInCircle(a, b, r, e, f) {
		return true
	} else {
		return false
	}
}

func inRect(a, b, x, y int) bool {
	if a > x || b > y || a < 0 || b < 0 {
		return false
	}
	return true
}

func liesInCircle(a, b, r int, e, f []int) bool {
	for i, _ := range e {
		//fmt.Println(e[i], f[i], a, b)
		if float64(r) >= math.Sqrt(math.Pow(float64(e[i]-a), 2)+math.Pow(float64(f[i]-b), 2)) {

			return true
		}
	}
	return false
}

package main

func f(s []int) []int {
	s = append(s, 9)
	return s
	//fmt.Println(s[0])
}

func f2() []int {
	var s []int
	s = append(s, 0)
	return s
}

func f3() []int {
	s := make([]int, 1)
	return s
}

func f4() {
	s := make([]int, 1) // why this does not escape to heap but f3 does
	_ = s
}

func main() {
	var s []int
	s = f(s)   // where does the array behind s lies? On heap then why didn't come in escape analysis
	s1 := f2() // Same issue.
	_ = s1
	//fmt.Println(s[0])

}

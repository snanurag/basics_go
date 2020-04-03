package main

func return8kb() [1024]int {
	var s [1024]int
	return s
}

func return8kbPt() *[1024]int {
	var s [1024]int
	return &s
}

func return80b() [10]int {
	var s [10]int
	return s
}

func return80bPt() *[10]int {
	var s [10]int
	return &s
}

func return8mb() [1024 * 1024]int {
	var s [1024 * 1024]int
	return s
}

func return8mbPt() *[1024 * 1024]int {
	var s [1024 * 1024]int
	return &s
}

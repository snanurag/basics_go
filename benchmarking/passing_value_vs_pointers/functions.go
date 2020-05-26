package main

func passByVal40B(v [5]int, i int) {
	for i < 10 {
		i++
		passByVal40B(v, i)
	}
}

func passByRef40B(v *[5]int, i int) {
	for i < 10 {
		i++
		passByRef40B(v, i)
	}
}

func passByVal80B(v [10]int, i int) {
	for i < 10 {
		i++
		passByVal80B(v, i)
	}
}

func passByRef80B(v *[10]int, i int) {
	for i < 10 {
		i++
		passByRef80B(v, i)
	}
}

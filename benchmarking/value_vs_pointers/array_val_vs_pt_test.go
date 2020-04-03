package main

import "testing"

var kb8 [1024]int
var kb8Pt *[1024]int

var b80 [10]int
var b80Pt *[10]int

var mb8 [1024 * 1024]int
var mb8Pt *[1024 * 1024]int

func BenchmarkCopying80b(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b80 = return80b()
	}
}

func BenchmarkCopying80bPt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b80Pt = return80bPt()
	}
}

func BenchmarkCopying8kb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		kb8 = return8kb()
	}
}

func BenchmarkCopying8kbPt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		kb8Pt = return8kbPt()
	}
}

func BenchmarkCopying64kb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mb8 = return8mb()
	}
}

func BenchmarkCopying64kbPt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mb8Pt = return8mbPt()
	}
}

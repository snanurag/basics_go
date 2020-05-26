package main

import "testing"

var arr40B [5]int
var arr80B [10]int

func BenchmarkPassByValue40B(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passByVal40B(arr40B, 0)
	}
}

func BenchmarkPassByRef40B(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passByRef40B(&arr40B, 0)
	}
}

func BenchmarkPassByValue80B(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passByVal80B(arr80B, 0)
	}
}

func BenchmarkPassByRef80B(b *testing.B) {
	for n := 0; n < b.N; n++ {
		passByRef80B(&arr80B, 0)
	}
}

package main

import (
	"testing"
)

func TestRead1(t *testing.T) {
	file := "./all.txt"
	Read1(file)
}

func TestRead2(t *testing.T) {
	file := "./all.txt"
	Read2(file)
}

func TestRead3(t *testing.T) {
	file := "./all.txt"
	Read3(file)
}

func BenchmarkRead1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		file := "./all.txt"
		Read1(file)
	}
}

func BenchmarkRead2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		file := "./all.txt"
		Read2(file)
	}
}

func BenchmarkRead3(b *testing.B) {

	for i := 0; i < b.N; i++ {
		file := "./all.txt"
		Read3(file)
	}
}

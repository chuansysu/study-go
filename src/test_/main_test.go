package main

import (
	"testing"
)

// go test -v -run TestFib main_test.go
// go test -v .
// go test -v -bench=. .
// go test -v -bench=Fib10
// go test -v -bench=Fib10 .
// go test -v -bench='.*Fib.*' .
// go test -v -bench=. -benchtime=3s -cpu=1,2,4 -count=3 -benchmem

// vscode运行-bench=.不行

func TestFib(t *testing.T) {
	fib(10)
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(20)
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(1, 2)
	}
}

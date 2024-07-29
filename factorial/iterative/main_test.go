package main

import (
	"testing"
)

//Unit test 
func TestFactorialIterative(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{5, 120},
		{10, 3628800},
	}

	for _, test := range tests {
		if result := factorialIterative(test.input); result != test.expected {
			t.Errorf("factorialIterative(%d) = %d, want: %d", test.input, result, test.expected)
		}
	}
}

//Benchmark test
func BenchmarkFactorialIterative(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		factorialIterative(999)
	}
}

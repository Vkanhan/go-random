package main

import (
	"testing"
)

func TestFactorialRecursion(t *testing.T) {
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
        if result := factorialRecursive(test.input); result != test.expected {
            t.Errorf("factorialRecursive(%d) = %d; want %d", test.input, result, test.expected)
        }
    }
}

func BenchmarkFactorialRecursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        factorialRecursive(999)
    }
}

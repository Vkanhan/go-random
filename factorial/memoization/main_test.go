package main

import (
	"testing"
)

func TestFactorialMemo(t *testing.T) {
	memo := Memo{1}
	
	testCases := []struct {
		input    uint
		expected uint
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{10, 3628800},
	}

	for _, tc := range testCases {
		result := memo.FactorialMemo(tc.input)
		if result != tc.expected {
			t.Errorf("Factorial(%d) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

func TestFactorialMemoization(t *testing.T) {
	memo := Memo{1}
	
	// First call to compute factorial of 5
	result1 := memo.FactorialMemo(5)
	if result1 != 120 {
		t.Errorf("Factorial(5) = %d; want 120", result1)
	}
	
	// Second call should use memoized value
	result2 := memo.FactorialMemo(5)
	if result2 != 120 {
		t.Errorf("Memoized Factorial(5) = %d; want 120", result2)
	}
	
	// Check if memo slice has correct length
	if len(memo) != 6 {
		t.Errorf("Memo length = %d; want 6", len(memo))
	}
}
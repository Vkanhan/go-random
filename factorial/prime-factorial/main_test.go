package main

import (
	"reflect"
	"testing"
)

// TestPrimeFactors tests the primeFactors function
func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		input    int
		expected map[int]int
	}{
		{1, map[int]int{}},               
		{2, map[int]int{2: 1}},          
		{12, map[int]int{2: 2, 3: 1}},    
		{15, map[int]int{3: 1, 5: 1}},   
	}

	// Iterate through all test cases
	for _, test := range tests {
		// Call the function we're testing
		result := primeFactors(test.input)
		// Compare the result to the expected output
		// DeepEqual is used because we're comparing maps
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("primeFactors(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// TestMergeFactors tests the mergeFactors function
func TestMergeFactors(t *testing.T) {
	// Define two input maps and the expected result of merging them
	a := map[int]int{2: 1, 3: 1}
	b := map[int]int{2: 2, 5: 1}
	expected := map[int]int{2: 3, 3: 1, 5: 1}

	// Call the function we're testing
	result := mergeFactors(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeFactors(%v, %v) = %v; want %v", a, b, result, expected)
	}
}

// TestFactorial tests the factorial function
func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int64
	}{
		{0, 1},         
		{1, 1},         
		{5, 120},       
		{10, 3628800},  
	}

	for _, test := range tests {
		// Call the function we're testing
		result := factorial(test.input)
		if result != test.expected {
			t.Errorf("factorial(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
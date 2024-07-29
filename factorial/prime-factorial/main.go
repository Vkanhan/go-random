package main

import (
	"fmt"
	"math"
)

// primeFactors calculates the prime factors of a given number n
// It returns a map where the key is the prime factor and the value is its exponent
func primeFactors(n int) map[int]int {
	factors := make(map[int]int)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}
	return factors
}

// mergeFactors combines two prime factor maps into a single map
// It's used to accumulate prime factors when calculating factorial
func mergeFactors(a, b map[int]int) map[int]int {
	result := make(map[int]int)
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] += v
	}
	return result
}

// factorial calculates the factorial of n using prime factorization
func factorial(n int) int64 {
	// Base case: factorial of 0 or 1 is 1
	if n <= 1 {
		return 1
	}

	// Accumulate prime factors for all numbers from 2 to n
	factors := make(map[int]int)
	for i := 2; i <= n; i++ {
		factors = mergeFactors(factors, primeFactors(i))
	}

	// Calculate the result by multiplying prime factors raised to their respective powers
	result := int64(1)
	for prime, exponent := range factors {
		result *= int64(math.Pow(float64(prime), float64(exponent)))
	}

	return result
}

func main() {
	// Calculate and print factorial for numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("Factorial of %d: %d\n", i, factorial(i))
	}
}
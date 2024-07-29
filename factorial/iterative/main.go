package main

import (
	"fmt"
)

//Find the factorial through iterative approach
func factorialIterative(n int) int {
	result := 1
	// Loop from 2 to n, multiplying the result by the loop variable i each time.
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Factorial of %d is %d\n", i, factorialIterative(i))
	}
}


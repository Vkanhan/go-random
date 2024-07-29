package main 

import(
	"fmt"
)

// factorialRecursive calculates the factorial of n using a recursive approach.
func factorialRecursive(n int) int {
    // Base case: 0! = 1.
    if n == 0 {
        return 1
    }
    // Recursive case: n! = n * (n-1)!.
    return n * factorialRecursive(n-1)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Factorial of %d is %d\n", i, factorialRecursive(i))
	}
}



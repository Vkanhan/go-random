package main

import "fmt"

// FactorialMemo is a type that represents a slice of uints to store the memoized factorial values.
type Memo []uint

// Factorial calculates the factorial of n using memoization.
// If the factorial value for n is not already computed, it extends the cache and computes the necessary values.
func (memo *Memo) FactorialMemo(n uint) uint {
	//dereference the memo pointer to access and manipulate the cache slice
    cache := *memo

    // Extend the cache until it includes the factorial value for n
    for uint(len(cache)) <= n {
        // Append the next factorial value to the cache
        cache = append(cache, uint(len(cache))*cache[len(cache)-1])
    }
    // Update the memo with the new cache to ensures that the changes are saved back to the original slice 
    *memo = cache
    // Return the factorial of n
    return cache[n]
}

func main() {
    // Initialize the memoization cache with the factorial of 1, which is 1
    memo := Memo{1}
    // Calculate and print the factorials from 1 to 10
    for i := 1; i <= 10; i++ {
        fmt.Printf("Factorial of %d is %d\n", i, memo.FactorialMemo(uint(i)))
    }
}


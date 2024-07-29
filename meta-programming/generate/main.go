package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define the template for the generated program
	template := `package factorial

import "fmt"

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	fmt.Println(Factorial(%d))
}
`

	// Define the number for which to generate the factorial program
	number := 5

	// Generate the program
	program := fmt.Sprintf(template, number)

	// Create a new directory for the generated program
	dir := "generated_program"
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the program to a file in the new directory
	filename := filepath.Join(dir, "main.go")
	err = os.WriteFile(filename, []byte(program), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Generated program written to %s\n", filename)
}
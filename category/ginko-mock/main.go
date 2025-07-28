package main

import (
	"fmt"
	"os"
)

type Spec struct {
	Description string
	TestFunc    func()
}

var currentSpecs []Spec

func Describe(description string, body func()) {
	fmt.Printf("Describe: %s\n", description)
	body()
	runSpecs()
}

func It(description string, testFunc func()) {
	spec := Spec{
		Description: description,
		TestFunc:    testFunc,
	}
	currentSpecs = append(currentSpecs, spec)
}

func runSpecs() {
	for _, spec := range currentSpecs {
		fmt.Printf("%s ", spec.Description)

		defer func() {
			if r := recover(); r != nil {
				fmt.Print("FAILED: \n", r)
				os.Exit(1)
			}
		}()
		spec.TestFunc()
		fmt.Println("PASSED")
	}
}

func AssertEqual(actual, expected any) {
	if actual != expected {
		panic(fmt.Sprintf("Expected %v but got %v", expected, actual))
	}
}

func main() {
	Describe("Adder", func() {
		It("Adds two numbers: ", func() {
			result := 2 + 3 
			AssertEqual(result, 5)
		})

		It("Fails when result is incorrect", func() {
			result := 2 + 2
			AssertEqual(result, 5)
		})
	})
}
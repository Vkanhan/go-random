// Package maps transform any function you want internally

package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func Map[In, Out any](in []In, fn func(In) Out) []Out {
	out := make([]Out, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func main() {
	//maps the slice of structs to extract its fields
	users := []User{
		{Name: "Einstein", Age: 24},
		{Name: "Tesla", Age: 26},
	}
	names := Map(users, func(u User) string {
		return u.Name
	})
	fmt.Println(names)

	ages := Map(users, func(u User) int {
		return u.Age
	})
	fmt.Println(ages)

	//maps the slice to square of it
	nums := []int{1, 2, 3, 4, 5}
	squares := Map(nums, func(x int) int {
		return x * x
	})
	fmt.Println(squares)

	//maps the slice to uppercase
	namess := []string{"india", "china"}
	uppercase := Map(namess, func(x string) string {
		return strings.ToUpper(x)
	})
	fmt.Println(uppercase)

	//maps the slice to length of its fields
	lengths := Map(namess, func(x string) int {
		return len(x)
	})
	fmt.Println(lengths)
}

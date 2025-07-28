package main

import (
	"fmt"
)

// category consistes of collection of objects can be types in programming
// collection of morphisms (arrows) between objects. can be functions in programming.
// Two rules -
// Identity - Each object has an identity morphism id: A -> A
// Associativity - If f: A -> B , g: B -> C, and h: C -> D then h * (g * f) = (h *g) * f

// Objects = Types

// Morphisms = Functions

// Category as an interface

// Composition is usual function composition

type Category[T any] interface {
	Id(x T) T
	Compose(f, g func(T) T) func(T) T
}

type IntCategory struct{}

func (IntCategory) Id(x int) int {
	return x
}

func (IntCategory) Compose(f, g func(x int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

type Morphism[A, B any] func(A) B

func Identity[T any]() Morphism[T, T] {
	return func(x T) T { return x }
}

func Composee[A, B, C any](f Morphism[B, C], g Morphism[A, B]) Morphism[A, C] {
	return func(x A) C {
		return f(g(x))
	}
}

func main() {
	// Identity law
	a := func(x int) int { return x + 1 }

	cat := IntCategory{}

	id := cat.Id

	a1 := cat.Compose(a, id)
	fmt.Println(a1(10))

	a2 := cat.Compose(id, a)
	fmt.Println(a2(10))

	// Associative law
	f := func(x int) int { return x + 1 }
	g := func(x int) int { return x * 2 }
	h := func(x int) int { return x - 3 }

	fg := cat.Compose(g, f)
	hfg := cat.Compose(h, fg)

	gh := cat.Compose(h, g)
	fgh := cat.Compose(gh, f)

	fmt.Println(hfg(5))
	fmt.Println(fgh(5))

	// Functions between different types
	m := func(x int) float64 { return float64(x) * 2.5 }
	n := func(s string) int { return len(s) }

	o := Composee(m, n) // h: string -> float64

	fmt.Println(o("abc"))
}

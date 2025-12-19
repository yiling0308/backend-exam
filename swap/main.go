package main

import (
	"fmt"
)

func swap[T any](a, b T) {
	npa, ok := any(a).(*int)
	if !ok {
		panic("swap[T] is not an *int")
	}
	npb, ok := any(b).(*int)
	if !ok {
		panic("swap[T] is not an *int")
	}

	*npa, *npb = *npb, *npa
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}

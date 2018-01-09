package main

import "fmt"

func test(x [2]int) {
	fmt.Printf("x: %p, %v\n", &x, x)
}

func main() {
	a := [...]int{
		10, 20,
	}

	var b [2]int
	b = a

	fmt.Printf("a: %p, %v\n", &a, a)
	fmt.Printf("b: %p, %v\n", &b, b)

	test(a)
}

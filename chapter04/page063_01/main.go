package main

import "fmt"

func test(x *int) {
	fmt.Printf("pointer: %p, target: %v\n", &x, x) // 形参
}

func main() {
	a := 0x100
	p := &a
	fmt.Printf("pointer: %p, target: %v\n", &p, p) // 实参

	test(p)
}

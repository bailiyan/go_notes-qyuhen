package main

import "fmt"

func main() {
	const v = 20
	fmt.Printf("%T, %v\n", v, v)

	var a byte = 10
	b := v + a
	fmt.Printf("%T, %v\n", b, b)

	const c float32 = 1.2
	d := c + v
	fmt.Printf("%T, %v\n", d, d)
}

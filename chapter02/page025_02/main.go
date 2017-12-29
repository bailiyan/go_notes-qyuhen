package main

import "fmt"

func main() {
	const (
		x uint16 = 120
		y
		s = "abc"
		z
	)

	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
}

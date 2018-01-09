package main

import "fmt"

func main() {
	x := [...]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	fmt.Printf("%T, %v\n", x[:], x[:])
	fmt.Printf("%T, %v\n", x[2:5], x[2:5])

}

package main

import "fmt"

func main() {
	var a [4]int

	b := [4]int{2, 5}
	c := [4]int{5, 3: 10}
	d := [...]int{1, 2, 3}
	e := [...]int{10, 3: 100}

	fmt.Println(a, b, c, d, e)
}

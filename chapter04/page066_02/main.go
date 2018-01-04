package main

import "fmt"

func test(a ...int) {
	fmt.Println(a)
	fmt.Printf("%T, %v\n", a, a)
}

func main() {
	a := [3]int{10, 20, 30}
	fmt.Printf("%T, %v\n", a, a)
	test(a[:]...)
}

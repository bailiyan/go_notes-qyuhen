package main

import "fmt"

func main() {
	b := make([]byte, 3)
	n := copy(b, "abcde")
	fmt.Println(n, b)
}

package main

import "fmt"

type N int

func main() {
	var n N = 5
	n.test()
	n.test1()
}

func (N) test() {
	fmt.Println("test...")
}

func (n N) test1() {
	fmt.Println(n)
}

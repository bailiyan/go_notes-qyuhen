package main

import "fmt"

type X struct{}

func (x *X) test() {
	fmt.Println("hi!", x)
}

func main() {
	var a *X
	a.test()

	//X{}.test()
}

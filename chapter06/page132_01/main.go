package main

import "fmt"

/*
 * 可以使用实例值或指针调用方法, 编译器会做处理
 */

type N int

func main() {
	var a N = 25
	p := &a

	a.value()
	a.pointer()

	p.value()
	p.pointer()
}

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

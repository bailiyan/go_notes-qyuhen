package main

import "unsafe"

func main() {
	const (
		ptrSize = unsafe.Sizeof(uintptr(0))
		strSize = len("hello, world")
	)

	println(ptrSize, strSize)
}

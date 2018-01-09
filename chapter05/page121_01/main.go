package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a struct{}
	var b [100]struct{}

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}

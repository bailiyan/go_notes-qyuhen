package main

import "fmt"

type flags byte

const (
	read  flags = 1 << iota
	write
	exec
)

func main() {
	println(read, write, exec)

	f := read | exec
	fmt.Printf("%b\n", f)
}

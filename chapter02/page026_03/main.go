package main

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	test(black)
	test(red)
	test(blue)
	test(100)

	// x:=2
	// test(x) // error
}

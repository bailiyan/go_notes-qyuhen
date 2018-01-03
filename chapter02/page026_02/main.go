package main

const (
	a = iota
	b
	c = 100
	d
	e = iota
	f
)

func main() {
	println(a, b, c, d, e, f)
}

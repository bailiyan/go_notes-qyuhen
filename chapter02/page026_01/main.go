package main

const (
	x = iota
	y
	z
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

const (
	_, _ = iota, iota * 10
	a, b
	c, d
)

func main() {
	println(x, y, z)
	println(KB, MB, GB)
	println(a, b, c, d)
}

package main

func main() {
	const x, y int = 123, 0x22
	const s = "hello, world!"
	const c = '我'

	println(x, y)
	println(s, c)

	const (
		i, f = 1, 0.123
		b    = false
	)
	println(i, f)
	println(b)
}

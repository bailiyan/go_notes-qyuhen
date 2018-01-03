package main

func main() {
	const x = 100
	const y byte = x
	println(x, y)

	const a int = 100
	// const b byte = a // error
}

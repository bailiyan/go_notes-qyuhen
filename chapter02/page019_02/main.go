package main

var x = 100

func main() {
	println(&x, x)

	x := 100
	println(&x, x)
}

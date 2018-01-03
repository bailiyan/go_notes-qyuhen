package main

func main() {
	a := 1
	p := &a
	*p++
	println(a)
}

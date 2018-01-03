package main

type data struct {
	x int
	s string
}

func main() {
	var a data = data{1, "abc"}
	println(a.x, a.s)

	b := data{
		2, "def"}
	println(b.x, b.s)

	c := []int{1, 2}
	for key, value := range c {
		println(key, value)
	}
}

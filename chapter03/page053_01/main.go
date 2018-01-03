package main

func count() int {
	print("count.")
	return 3
}

func main() {
	for i, c := 0, count(); i < c; i++ {
		println("a", i)
	}

	c := 0
	for c < count() {
		println("b", c)
		c++
	}
}

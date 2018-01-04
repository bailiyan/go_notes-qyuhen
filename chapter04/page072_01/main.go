package main

func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	f := test(123)
	f()
	g := test(321)
	g()
}

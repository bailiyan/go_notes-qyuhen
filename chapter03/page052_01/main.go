package main

func main() {
	switch x := 5; {
	case x > 5:
		println("a")
	case x > 0 && x <= 4:
		println("b")
	default:
		println("z")
	}
}

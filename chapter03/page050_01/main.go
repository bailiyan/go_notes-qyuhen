package main

func main() {
	switch x := 6; x {
	case 5:
		x += 50
		println(x)
	default:
		x += 100
		println(x)
	}
}

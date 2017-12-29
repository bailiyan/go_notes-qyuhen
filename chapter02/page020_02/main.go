package main

func main() {
	x := 100
	println(&x, x)

	//{
		x, y := 200, 300
		println(&x, x, y)
	//}
}

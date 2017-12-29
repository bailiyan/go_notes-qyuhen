package main

func main() {
	const x = 123
	println(x)

	const y = 1.23

	{
		const x = "abc"
		println(x)
	}

	println(x)
}

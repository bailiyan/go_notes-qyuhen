package main

func main() {
	data := [3]string{"a", "b", "c"}

	for i := range data {
		println(i, data[i])
	}

	for _, s := range data {
		println(s)
	}

	for range data {

	}
}

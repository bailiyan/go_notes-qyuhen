package main

func main() {
	x := []int{100, 101, 102}

	for k, v := range x {
		println(k, ":", v)
	}

}

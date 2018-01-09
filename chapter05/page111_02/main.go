package main

func main() {
	m := make(map[string]int)

	for i := 0; i < 8; i++ {
		m[string('a'+i)] = i
	}

	for i := 0; i < 4; i++ {
		for k, v := range m {
			print(k, ":", v, " ")
		}

		println()
	}

}

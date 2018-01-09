package main

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	m["a"] = 10
	m["c"] = 30

	if v, ok := m["d"]; ok {
		println(v)
	}

	delete(m, "d")
}

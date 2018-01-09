package main

import "fmt"

func main() {
	type user struct {
		name string
		age  byte
	}

	m := map[int]user{
		1: {"Tom", 19},
	}

	u := m[1]
	u.age += 1
	m[1] = u
	fmt.Println(m[1])

	m2 := map[int]*user{
		1: &user{"Jack", 20},
	}

	m2[1].age++
	fmt.Println(*m2[1])
}

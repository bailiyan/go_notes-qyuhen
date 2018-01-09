package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	p := &user{
		name: "Tom",
		age:  20,
	}

	p.name = "Marry"
	p.age++

	fmt.Println(*p)
}

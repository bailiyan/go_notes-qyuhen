package main

import "fmt"

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	tile string
}

func main() {
	var m manager

	m.name = "Tom"
	m.age = 29
	m.tile = "CTO"

	fmt.Println(m)
}

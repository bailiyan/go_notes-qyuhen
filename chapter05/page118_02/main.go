package main

import "fmt"

type user struct {
	name string
	age  byte
}

func main() {
	u1 := user{"Tom", 12}
	u2 := user{name: "Tom"}
	fmt.Println(u1, u2)

}

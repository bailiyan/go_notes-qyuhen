package main

import "fmt"

type user struct{}

type manager struct {
	user // 匿名字段
}

func (user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + "; manager"
}

func main() {
	var m manager

	fmt.Println(m.toString())
	fmt.Println(m.user.toString())
}

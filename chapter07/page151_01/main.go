package main

import "fmt"

func main() {
	// 将一个函数赋值给空接口
	var x interface{} = func(x int) string {
		return fmt.Sprintf("d:%d", x)
	}

	// type switch
	switch v := x.(type) {
	case nil:
		fmt.Println("nil")
	case *int:
		fmt.Println("*v")
	case fmt.Stringer:
		fmt.Println(v)
	case func(int) string:
		fmt.Println(v(100))
	default:
		fmt.Println("unknown")
	}
}

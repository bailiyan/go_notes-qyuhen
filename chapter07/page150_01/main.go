package main

import "fmt"

type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

func main() {
	var d data = 15
	var x interface{} = d

	/*
	 * 类型断言
	 */

	// 1. 转换为更具体的接口类型
	if n, ok := x.(fmt.Stringer); ok {
		fmt.Println(n)
	}

	// 2. 转换回原始类型
	if d2, ok := x.(data); ok {
		fmt.Println(d2)
	}

}

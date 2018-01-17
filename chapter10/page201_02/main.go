package main

import (
	"fmt"
	"reflect"
)

/*
 * 通过反射构造类型
 */
func main() {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))

	fmt.Println(a, m)
}

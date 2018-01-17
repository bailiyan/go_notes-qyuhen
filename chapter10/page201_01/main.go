package main

import (
	"reflect"
	"fmt"
)

type X int
type Y int

func main() {
	var a, b X = 100, 200
	var c Y = 300

	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)

	fmt.Println(ta.Name(), tb.Name(), tc.Name())
	fmt.Println(ta.Kind(), tb.Kind(), tc.Kind())

}

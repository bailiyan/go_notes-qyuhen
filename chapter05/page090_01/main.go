package main

import "fmt"

func main() {
	var bs []byte
	bs = append(bs, "abc"...)

	fmt.Println(bs)
}

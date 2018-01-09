package main

import "unicode/utf8"

func main() {
	s := "雨.痕"
	println(len(s), utf8.RuneCountInString(s))
}

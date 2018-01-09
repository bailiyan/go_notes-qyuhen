package main

func main() {
	s := "ab" +
		"cd"

	println(s == "abcd")
	println(s > "abc")

	s = "abc"
	println(s[1])
	println(string((s[1])))
	println(&s[1])
}

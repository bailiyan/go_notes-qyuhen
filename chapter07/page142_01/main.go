package main

/*
 * 编译器根据方法集来判断是否实现了某接口
 */
type tester interface {
	test()
	string() string
}

type data struct{}

func (*data) test()         {}
func (data) string() string { return "12345" }

func main() {
	var d data

	// var t tester = d // panic

	var t tester = &d
	t.test()
	println(t.string())
}

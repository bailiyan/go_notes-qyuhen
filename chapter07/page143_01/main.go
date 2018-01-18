package main

type stringer interface {
	string() string
}

type tester interface {
	stringer // 嵌入接口
	test()
}

type data struct{}

func (*data) test() {}

func (data) string() string {
	return "123"
}

func main() {
	// data 的方法集是 test()
	// *data 的方法集是 test(), string()
	// tester 接口声明的方法包括 test() string()
	// 综上所述， *data 类型实现 tester 接口
	var d data
	var t tester = &d
	t.test()
	println(t.string())
}

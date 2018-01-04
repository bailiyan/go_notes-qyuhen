package main

import "errors"

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func main() {
	r, e := div(100, 0)
	if e != nil {
		println("error!")
	} else {
		println(r)
	}

}

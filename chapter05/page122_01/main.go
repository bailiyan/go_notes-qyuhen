package main

import "fmt"

func main() {
	exit := make(chan struct{})

	go func() {
		fmt.Println("hello, world!")
		exit <- struct{}{}
	}()

	<-exit
	fmt.Println("end.")
}

package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {
	// 最後に実行される
	defer fmt.Println("world")
	foo()
	fmt.Println("hello")

	// stacking defer
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
}

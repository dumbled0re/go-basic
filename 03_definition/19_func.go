package main

import "fmt"

// func add(x int, y int) (int, int) {
func add(x, y int) (int, int) {
	return x + y, x - y
}

// 返り値の変数を指定できる(その場合は変数名が分かっているので、returnだけでもok)
func cal(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	// 変数に関数を代入可能
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	// 関数をそのまま実行
	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}

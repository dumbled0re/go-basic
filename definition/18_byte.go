package main

import (
	"fmt"
)

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	// 文字列にキャスト
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}

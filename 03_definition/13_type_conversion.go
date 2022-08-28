package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数値のキャスト
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	// 文字列のキャスト
	var s string = "14"
	// 文字列とエラーを返す
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	// アスキーコードを取得
	fmt.Println(h[0])
	// 文字列を取得
	fmt.Println(string(h[0]))
}

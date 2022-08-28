package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	// Hello WorldのHを取得
	fmt.Println(string("Hello World"[0]))
	var s string = "Hello World"
	// 文字列の置換
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	// sにWorldという文字列が含まれていたらtrueを返す
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(`Test
						Test
Test`)
	// 特殊文字は\で表示(``でもok)
	fmt.Println("\"")
	fmt.Println(`"`)
}

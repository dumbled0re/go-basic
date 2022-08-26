package main

import "fmt"

func main() {
	// ディクショナリ
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	// 値を変更
	m["banana"] = 300
	fmt.Println(m)
	// 新しい要素追加
	m["new"] = 500
	fmt.Println(m)

	// 有れば100、無ければ0
	fmt.Println(m["nothing"])

	// okにはtrueかfalse
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// メモリに空のマップを作成
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	/*
		宣言しているだけで、メモリに保存されていなのでエラーが発生(マップのみ)
		var m3 map[string]int
		m3["pc"] = 5000
		fmt.Println(m3)
	*/

	// 宣言したマップとスライスはnil判定される
	var m3 map[string]int
	if m3 == nil {
		fmt.Println("Nil")
	}
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}

package main

import "fmt"

func main() {

	// タイプがポインターの場合はnewを使い、それ以外はmakeを使う
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p *int = new(int)
	fmt.Printf("%T\n", p)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	/*
		// メモリにポインター型を作成
		var p *int = new(int)
		fmt.Println(*p) // 0xc0000160d8
		*p++
		fmt.Println(*p)

		// 宣言しただけでメモリに確保はしていない
		var p2 *int
		fmt.Println(p2) // <nil>
		*p2++
		fmt.Println(p2)
	*/

}

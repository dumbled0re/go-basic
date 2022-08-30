package main

import "fmt"

// 変数名は大文字で記載する必要があり
type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
	// fmt.Println(v)
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(*v2)

	/*
		v := Vertex{X: 1, Y: 2}
		fmt.Println(v)
		fmt.Println(v.X, v.Y)
		v.X = 100
		fmt.Println(v.X, v.Y)

		v2 := Vertex{X: 1}
		fmt.Println(v2)

		v3 := Vertex{1, 2, "test"}
		fmt.Println(v3)

		// v4とv5は同じ意味
		v4 := Vertex{}
		fmt.Printf("%T %v\n", v4, v4)

		var v5 Vertex
		fmt.Printf("%T %v\n", v5, v5)

		// v6とv7は同じ意味だけどv7の方が明示的にポインターが返ってくるのが分かる
		v6 := new(Vertex)
		fmt.Printf("%T %v\n", v6, v6)

		// newよりこちらを推奨
		v7 := &Vertex{}
		fmt.Printf("%T %v\n", v7, v7)

		// sliceとmapはmakeを使った初期化が推奨
		s := make([]int, 0)
		// s := []int{}
		fmt.Println(s)
	*/
}

package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// rangeを使った書き方(indexとvalueが来る)
	for i, v := range l {
		fmt.Println(i, v)
	}

	// 値だけ取り出す
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// keyだけ取り出す場合はvを省略可能
	for k := range m {
		fmt.Println(k)
	}

	// valueだけ取り出したい場合は_を使わないといけない
	for _, v := range m {
		fmt.Println(v)
	}
}

package main

import "fmt"

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	// short variable declarations(関数内でしか使えない)
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	// タイプと値を出力
	fmt.Printf("type=%T value=%v\n", xf64, xf64)
	fmt.Printf("type=%T value=%v\n", xi, xi)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}

package main

import "fmt"

func main() {
	// Q1
	var i int = 10
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j) // 10

	// Q2
	var i2 int = 100
	var j2 int = 200
	var p1 *int
	var p2 *int
	p1 = &i2
	p2 = &j2
	i2 = *p1 + *p2
	p2 = p1
	j2 = *p2 + i2   // 100 + 300
	fmt.Println(j2) // 400
}

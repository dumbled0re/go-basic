package main

import "fmt"

func main() {
	// Q1
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var min int
	for i, v := range l {
		if i == 0 {
			min = v
			continue
		}
		if min > v {
			min = v
		}
	}
	fmt.Println("Q1の答え:", min)

	// Q2
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	total_sum := 0
	for _, v := range m {
		total_sum += v
	}

	fmt.Println("Q2の答え:", total_sum)
}

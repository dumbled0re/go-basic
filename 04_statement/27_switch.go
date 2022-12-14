package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	// os := getOsName()
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	// どれにも当てはまらない時
	default:
		fmt.Println("Default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	// 条件を省略可能
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}

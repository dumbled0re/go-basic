package main

import (
	"fmt"
	"os/user"
	"time"
)

// main関数より先に呼ばれる
func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	// bazz()
	fmt.Println("Hello World!", time.Now())
	fmt.Println(user.Current())
}

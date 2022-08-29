package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./30_error_handling.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	// 1つ以上新しく変数を宣言するならshort declarationを使える
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))
}

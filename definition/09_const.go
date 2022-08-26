package main

import "fmt"

// 書き換え不要変数
// 型宣言不要
const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)
}

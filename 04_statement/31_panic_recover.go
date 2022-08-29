package main

import "fmt"

func thirdPartyConnectDB() {
	// 例外を発生させる(GOでは推奨されていない)
	panic("Unable to connect database!")
}

func save() {
	// recoverでエラーが発生した時にキャッチする
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}

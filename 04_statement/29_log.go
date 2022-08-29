package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 外部に公開する場合はアッパーキャメルケース（先頭大文字から始まる)
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("logs/test.log")
	_, err := os.Open("sdafaggas")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	// エラーが発生してプログラム終了
	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!!")

	fmt.Println("ok")
}

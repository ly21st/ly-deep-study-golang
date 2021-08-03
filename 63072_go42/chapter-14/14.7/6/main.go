package main

import (
	"log"
	"os"
)

func main() {
	LogFile, err := os.OpenFile("./My.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	//若果文件不存在就创建文件并打开
	defer LogFile.Close()

	if err != nil {
		log.Fatalln("fail to create My.log file!")
	}

	logger := log.New(LogFile, "[info]", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("Log info")

	logger.SetPrefix("[debug]")
	logger.Println("Log Debug")
}

package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	// 把函数 timeHandler 转为 HandlerFunc
	th := http.HandlerFunc(timeHandler)
	// th作为mux的handler
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

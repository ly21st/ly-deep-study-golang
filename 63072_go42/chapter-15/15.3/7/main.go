package main

import (
	"fmt"
	"net/http"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", myfunc)
	http.ListenAndServe(":8080", mux)
}

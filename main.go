package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server is running...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":3000", nil)
}

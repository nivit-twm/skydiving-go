package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":80", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

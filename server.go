package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "BREW" {
		fmt.Fprintf(w, "Brewing tea.")
	} else {
		http.Error(w, "The requested entity body is short and stout.", http.StatusTeapot)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3011", nil)
}

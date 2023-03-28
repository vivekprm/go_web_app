package main

import "net/http"

func main() {
	http.HandleFunc("/", someFunc)
	// nil means we going to use default ServeMux
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Universe!"))
}
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/terms", terms)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello from Home Page!`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello from About Page!`)
}

func terms(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello from Terms Page!`)
}

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
	fmt.Fprintf(w, "<p>Go is awesome!!!</p>")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8000", nil)

}

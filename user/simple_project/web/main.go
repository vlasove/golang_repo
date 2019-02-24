package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Golang is awesome!</h1>")
	fmt.Fprintf(w, "<p>Go fast and simple</p>")
	fmt.Fprintf(w, "<p>You can %v even add %v </p>", "pidor", "<strong>variables</strong>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Evgen")
}

func main() {
	http.HandleFunc("/", index_handler)
	//http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":3000", nil)

}

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Evgen"
	}
	fmt.Fprint(res, "Haudy, My name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Evgen"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	fmt.Fprint(res, "This is the homepage")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.ListenAndServe("localhost:8000", nil)
}

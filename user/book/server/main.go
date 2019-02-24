package main

import (
	"fmt"
	"net/http"
	"os"
)

func kill(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Stop server")
	os.Exit(0)
}

func hello(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "<h1>Smart WEB</h1>")
	fmt.Fprint(res, "<p>by Evgen</p>")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/kill", kill)

	http.ListenAndServe("localhost:8000", nil)
}

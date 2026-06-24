package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ASCII Art Web!")
}
func main() {
	// created a route to request a path to homeHandler
	http.HandleFunc("/", homeHandler)

	fmt.Print("Server Starting on :8080 ")
	http.ListenAndServe(":8080", nil)
}

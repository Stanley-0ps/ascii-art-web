package main

import (
	"ascii-art/functions"
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html") // loading html content from html file

	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	err = template.Execute(w, nil) // send Html content back to browser

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	fileName := "banners/" + banner + ".txt"

	asciiMap, err := functions.BuildFontMap(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result, err := functions.RenderAscii(text, asciiMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	data := PageData{
		Result: result,
	}

	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func main() {
	// registered a route to request a path to homeHandler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Print("Server Starting on :8080 ")
	http.ListenAndServe(":8080", nil)
	if ere!=nil{
		fmt.
	}
}

package main

import (
	"fmt"
	//"log"
	"html/template"
	"net/http"
)

func handlerHttpTemplates(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, error := template.ParseFiles("templates/index.html")
	if error != nil {
		fmt.Fprintf(w, error.Error())
	}

	htmlTemplate.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.HandleFunc("/", handlerHttpTemplates)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"

	"go-google-calendar-statistic/calendar"
)

func handlerHttpTemplates(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, error := template.ParseFiles("templates/index.html")
	if error != nil {
		fmt.Fprintf(w, error.Error())
	}

	htmlTemplate.ExecuteTemplate(w, "index", nil)
}

func main() {
	calendar.Auth()

	http.HandleFunc("/", handlerHttpTemplates)
	http.ListenAndServe(":8080", nil)

}

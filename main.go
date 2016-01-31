package main

import (
	"fmt"
	"go-google-calendar-statistic/myCalendar"
	"html/template"
	"log"
	"net/http"

	//"google.golang.org/api/calendar/v3"
)

func handlerHttpTemplates(w http.ResponseWriter, r *http.Request) {
	htmlTemplate, error := template.ParseFiles("templates/index.html")
	if error != nil {
		fmt.Fprintf(w, error.Error())
	}

	htmlTemplate.ExecuteTemplate(w, "index", nil)
}

func logError(err error) {
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
}

func main() {
	svr, err := myCalendar.Auth()
	logError(err)

	events, err := myCalendar.GetNextTenEvents(svr)
	logError(err)

	myCalendar.PrintEvents(events, "Upcoming 10 events")

	oldEvents, err := myCalendar.GetLastMonthEvents(svr)
	logError(err)

	myCalendar.PrintEvents(oldEvents, "Last month events:")

	//http.HandleFunc("/", handlerHttpTemplates)
	//http.ListenAndServe(":8080", nil)

}

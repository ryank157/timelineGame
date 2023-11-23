package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"text/template"
)

type Film struct {
	Event string
	Year int64
}

func main() {
	events := map[string][]Film{
		"Events":{
			{Event: "The Godfather", Year: 2000},
			{Event: "The Godfather", Year: 1998},
			{Event: "The Godfather", Year: 1999},
			{Event: "The Godfather", Year: 2001},
		},
	}
	sort.Slice(events["Events"], func(i,j int) bool {
		return events["Events"][i].Year < events["Events"][j].Year
	})

	h1 := func (w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("./src/index.html"))
		

		tmpl.Execute(w, events)
	} 

	addEvent := func (w http.ResponseWriter, r *http.Request){
		fmt.Println("made it here")
		if err := r.ParseForm(); err != nil {
			fmt.Println("unknown error")
			// Handle error
			return
		}
	
		eventDescription := r.FormValue("eventDescription")
		eventYearStr := r.FormValue("eventYear")
		eventYear, err := strconv.ParseInt(eventYearStr, 10, 64)
		if err != nil {
			fmt.Println("unknown error 2")
			// Handle error
			// Handle error
			return
		}

	
	   // Create a new Film instance
	   newEvent := Film{
        Event: eventDescription,
        Year: eventYear,
    }

    // Append the new event to the Events slice
    events["Events"] = append(events["Events"], newEvent)
	fmt.Print(events)

    // Re-sort the Events slice
    sort.Slice(events["Events"], func(i, j int) bool {
        return events["Events"][i].Year < events["Events"][j].Year
    })

    tmpl := template.Must(template.ParseFiles("./src/index.html"))
    tmpl.Execute(w, events)
	}
	http.HandleFunc("/",h1)
	http.HandleFunc("/addEvent",addEvent)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"text/template"
)

type Film struct {
	Event string
	Year int64
}

func main() {
	fmt.Println("hello world")

	h1 := func (w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("./src/index.html"))
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

		tmpl.Execute(w, events)
	}
	http.HandleFunc("/",h1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
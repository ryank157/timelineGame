package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"text/template"
)

type Film struct {
	Event string
	Year  int64
}

func readDB() (records [][]string) {

	// Open the CSV file
	csvFile, err := os.Open("events.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()

	// Read the CSV file
	reader := csv.NewReader(csvFile)
	records, err = reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("brian")
	// Print the records
	for _, record := range records {
		fmt.Println("first")
		fmt.Println(record[0])
		fmt.Println("second")
		fmt.Println(record[1])
	}

	return
}

func main() {
	events2 := readDB()
	fmt.Print(events2)
	events := map[string][]Film{
		"Events": {
			{Event: "The Godfather", Year: 2000},
			{Event: "The Godfather", Year: 1998},
			{Event: "The Godfather", Year: 1999},
			{Event: "The Godfather", Year: 2001},
		},
	}

	sort.Slice(events["Events"], func(i, j int) bool {
		return events["Events"][i].Year < events["Events"][j].Year
	})

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./src/index.html", "./src/room.html"))

		err := tmpl.Execute(w, events)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			// Optionally send an error response to the client
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
			return
		}

	}

	addEvent := func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("made it here")
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

		tmpl := template.Must(template.ParseFiles("./src/index.html"))

		tmpl.ExecuteTemplate(w, "eventListElement", Film{
			Event: eventDescription,
			Year:  eventYear,
		})

	}
	// joinRoom := func (w http.ResponseWriter, r *http.Request){
	// 	tmpl := template.Must(template.ParseFiles("./src/room.html"))
	// }
	http.HandleFunc("/", h1)
	http.HandleFunc("/addEvent", addEvent)
	// http.HandleFunc("/joinRoom",joinRoom)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

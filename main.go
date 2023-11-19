package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title string
	Director string
}

func main() {
	fmt.Println("hello world")

	h1 := func (w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("./src/index.html"))
		films := map[string][]Film{
			"Films":{
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
			},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/",h1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
package main

import (
	"html/template"
	"net/http"
)

func main() {
	tpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tpl.Execute(writer, nil)
	})
	http.ListenAndServe(":8091", nil)
}

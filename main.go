package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Member struct {
	Id           int
	Name         string
	Email        string
	RegisterDate string
}

type Members struct {
	List []Member
}

func main() {

	data := Members{List: []Member{
		{
			Id:           3,
			Name:         "Member3",
			Email:        "mail@com.com",
			RegisterDate: "15.04.2000",
		},
	}}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		name := request.FormValue("Name")
		fmt.Println(name)
		tpl, _ := template.ParseFiles("index.html")
		tpl.Execute(writer, data)
	})
	http.ListenAndServe(":8091", nil)
}

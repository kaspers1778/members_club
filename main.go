package main

import (
	"html/template"
	"net/http"
	"time"
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

func (ms *Members) Add(m Member) {
	if !ms.Contatin(m) {
		ms.List = append(ms.List, m)
	}
}

func (ms *Members) Contatin(m Member) bool {
	for _, el := range ms.List {
		if el.Email == m.Email {
			return true
		}
	}
	return false
}

func (ms *Members) Clear() {
	ms.List = []Member{}
}

func main() {
	members := Members{List: []Member{}}
	http.HandleFunc("/", makeSlashHandler(&members))
	http.HandleFunc("/add", makeAddHandler(&members))
	http.HandleFunc("/clear", makeClearHandler(&members))

	http.ListenAndServe(":8091", nil)
}

func makeSlashHandler(members *Members) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("index.html")
		tpl.Execute(w, members)
	}
}

func makeAddHandler(members *Members) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newMember := Member{
			Id:           len(members.List) + 1,
			Name:         r.FormValue("Name"),
			Email:        r.FormValue("Email"),
			RegisterDate: time.Now().Format("02.01.2006"),
		}
		members.Add(newMember)
		tpl, _ := template.ParseFiles("index.html")
		tpl.Execute(w, members)
	}
}

func makeClearHandler(members *Members) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		members.Clear()
		tpl, _ := template.ParseFiles("index.html")
		tpl.Execute(w, members)
	}
}

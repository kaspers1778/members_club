package main

import (
	"fmt"
	"html/template"
	"members_club/internal"
	"net/http"
	"os"
	"time"
)

func main() {
	PORT := "8000"
	if v, ok := os.LookupEnv("PORT"); ok {
		PORT = v
	}
	message := internal.Message{
		Members: []internal.Member{},
		Errors:  map[string]string{},
	}
	var sl internal.ServerLogger
	sl.New()
	sl.Info(fmt.Sprintf("Server has been started on :%v ", PORT))
	http.HandleFunc("/", makeSlashHandler(&message))
	http.HandleFunc("/add", makeAddHandler(&message, sl))
	http.HandleFunc("/clear", makeClearHandler(&message, sl))

	http.ListenAndServe(":"+PORT, nil)
}

func makeSlashHandler(members *internal.Message) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("index.html")
		tpl.Execute(w, members)
	}
}

func makeAddHandler(message *internal.Message, log internal.ServerLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		message.ClearErrors()
		newMember := internal.Member{
			Id:           len(message.Members) + 1,
			Name:         r.FormValue("Name"),
			Email:        r.FormValue("Email"),
			RegisterDate: time.Now().Format("02.01.2006"),
		}
		log.LogAddRequest(newMember)
		message.Add(newMember)
		log.LogAddResponse(message.Errors)
		http.Redirect(w, r, "", http.StatusSeeOther)

	}
}

func makeClearHandler(members *internal.Message, log internal.ServerLogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.LogClearRequest()
		members.Clear()
		log.LogClearResponse()
		http.Redirect(w, r, "", http.StatusSeeOther)
	}
}

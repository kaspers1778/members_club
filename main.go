package main

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
	"time"
)

type Member struct {
	Id           int
	Name         string
	Email        string
	RegisterDate string
}

func (m Member) Validate() (nameErr, emailErr error) {
	isNameValid, _ := regexp.MatchString("[\\w\\.\\s]+", m.Name)
	if !isNameValid {
		nameErr = errors.New("Name should be valid")
	}
	isEmailValid, _ := regexp.MatchString("[\\w\\.\\d]+@[\\w\\d]+\\.[\\w]+", m.Email)
	if !isEmailValid {
		emailErr = errors.New("Email should be valid")
	}
	return nameErr, emailErr

}

type Message struct {
	Members []Member
	Errors  map[string]string
}

func (ms *Message) ClearErrors() {
	ms.Errors = map[string]string{}
}

func (ms *Message) ValidateMember(m Member) bool {
	nameErr, emailErr := m.Validate()
	if nameErr != nil {
		ms.Errors["Name"] = nameErr.Error()
	}
	if emailErr != nil {
		ms.Errors["Email"] = emailErr.Error()
	}

	if nameErr == nil && emailErr == nil {
		return true
	}
	return false
}

func (ms *Message) Add(m Member) {
	isMemberValid := ms.ValidateMember(m)
	if ms.Contain(m) {
		ms.Errors["Email"] = "Members club already has member with such email"
	} else if isMemberValid {
		ms.Members = append(ms.Members, m)
	}
}

func (ms *Message) Contain(m Member) bool {
	for _, el := range ms.Members {
		if el.Email == m.Email {
			return true
		}
	}
	return false
}

func (ms *Message) Clear() {
	ms.Members = []Member{}
}

func main() {
	message := Message{Members: []Member{},
		Errors: map[string]string{}}
	http.HandleFunc("/", makeSlashHandler(&message))
	http.HandleFunc("/add", makeAddHandler(&message))
	http.HandleFunc("/clear", makeClearHandler(&message))

	http.ListenAndServe(":8091", nil)
}

func makeSlashHandler(members *Message) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("index.html")
		tpl.Execute(w, members)
	}
}

func makeAddHandler(message *Message) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		message.ClearErrors()
		newMember := Member{
			Id:           len(message.Members) + 1,
			Name:         r.FormValue("Name"),
			Email:        r.FormValue("Email"),
			RegisterDate: time.Now().Format("02.01.2006"),
		}
		message.Add(newMember)
		http.Redirect(w, r, "", http.StatusSeeOther)

	}
}

func makeClearHandler(members *Message) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		members.Clear()
		http.Redirect(w, r, "", http.StatusSeeOther)
	}
}

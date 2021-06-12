package internal

import (
	"errors"
	"regexp"
)

type Member struct {
	Id           int
	Name         string
	Email        string
	RegisterDate string
}

func (m Member) Validate() (nameErr, emailErr error) {
	isNameValid, _ := regexp.MatchString("[a-zA-Z\\.\\s]+$", m.Name)
	if !isNameValid {
		nameErr = errors.New("Name should be valid")
	}
	isEmailValid, _ := regexp.MatchString("[\\w\\.\\d]+@[\\w\\d]+\\.[\\w]+", m.Email)
	if !isEmailValid {
		emailErr = errors.New("Email should be valid")
	}
	return nameErr, emailErr

}

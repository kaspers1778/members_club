package tests

import (
	"members_club/internal"
	"testing"
)

var m = internal.Message{
	Members: []internal.Member{},
	Errors:  map[string]string{},
}

func TestMessageAddMemberRight(t *testing.T) {
	validMember := internal.Member{
		Id:           1,
		Name:         "John",
		Email:        "john@gmail.com",
		RegisterDate: "01.02.1999",
	}
	m.Add(validMember)
	if len(m.Errors) != 0 {
		t.Errorf("message should not have any errors, actual %v", len(m.Errors))
	}
	m.ClearErrors()
}

func TestMessageAddMemberWrong(t *testing.T) {
	validMember := internal.Member{
		Id:           1,
		Name:         "John564",
		Email:        "john@@@gmail.com",
		RegisterDate: "01.02.1999",
	}
	m.Add(validMember)
	if len(m.Errors) != 2 {
		t.Errorf("message should have 2 errors, actual %v", len(m.Errors))
	}
	m.ClearErrors()
}

func TestMessageAddMemberRepeat(t *testing.T) {
	validMember := internal.Member{
		Id:           1,
		Name:         "John",
		Email:        "john@gmail.com",
		RegisterDate: "01.02.1999",
	}
	m.Add(validMember)
	if len(m.Errors) != 1 {
		t.Errorf("message should have 1 error, actual %v", len(m.Errors))
	}
	m.ClearErrors()
}

func TestMessageClearErrors(t *testing.T) {
	validMember := internal.Member{
		Id:           1,
		Name:         "John564",
		Email:        "john@@@gmail.com",
		RegisterDate: "01.02.1999",
	}
	m.Add(validMember)
	if len(m.Errors) != 2 {
		t.Errorf("message should have 2 errors, actual %v", len(m.Errors))
	}
	m.ClearErrors()
	if len(m.Errors) != 0 {
		t.Errorf("message should have not any errors, actual %v", len(m.Errors))
	}
}

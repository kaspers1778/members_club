package tests

import (
	"members_club/internal"
	"testing"
)

func TestMemberValidateRight(t *testing.T) {
	validMember := internal.Member{
		Id:           1,
		Name:         "John",
		Email:        "john@gmail.com",
		RegisterDate: "01.02.1999",
	}
	nameErr, emailErr := validMember.Validate()
	if nameErr != nil {
		t.Errorf("nameErr must be nil")
	}
	if emailErr != nil {
		t.Errorf("emailErr must be nil")
	}
}

func TestMemberValidateWrongName(t *testing.T) {
	validMember := internal.Member{
		Id:           1,
		Name:         "John54684",
		Email:        "john@gmail.com",
		RegisterDate: "01.02.1999",
	}
	nameErr, emailErr := validMember.Validate()
	if nameErr == nil {
		t.Errorf("nameErr should not be nil")
	}
	if emailErr != nil {
		t.Errorf("emailErr must be nil")
	}
}

func TestMemberValidateWrongEmail(t *testing.T) {
	validMember := internal.Member{
		Id:           1,
		Name:         "John",
		Email:        "john@@45612.ua.com",
		RegisterDate: "01.02.1999",
	}
	nameErr, emailErr := validMember.Validate()
	if nameErr != nil {
		t.Errorf("nameErr should be nil")
	}
	if emailErr == nil {
		t.Errorf("emailErr should not nil")
	}
}

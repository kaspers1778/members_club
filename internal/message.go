package internal

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

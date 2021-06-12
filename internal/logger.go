package internal

import (
	"fmt"
	"log"
	"os"
)

type ServerLogger struct {
	logger *log.Logger
}

func (l *ServerLogger) New() {
	l.logger = log.New(os.Stdout, "", log.LstdFlags)
}

func (l *ServerLogger) Info(m string) {
	l.logger.Println(m)
}

func (l *ServerLogger) LogAddRequest(m Member) {
	l.logger.Println(fmt.Sprintf("Request - name: %v, email: %v.", m.Name, m.Email))
}
func (l *ServerLogger) LogAddResponse(errors map[string]string) {
	if len(errors) == 0 {
		l.logger.Println("Response - OK.")
	} else {
		errorsStr := ""
		for _, err := range errors {
			errorsStr += fmt.Sprintf("%v, ", err)
		}
		if errorsStr != "" {
			errorsStr = errorsStr[:len(errorsStr)-2]
		}
		l.logger.Println(fmt.Sprintf("Response - %v.", errorsStr))
	}
}

func (l *ServerLogger) LogClearRequest() {
	l.logger.Println(fmt.Sprintf("Request - clear members."))
}
func (l *ServerLogger) LogClearResponse() {
	l.logger.Println(fmt.Sprintf("Response - OK."))
}

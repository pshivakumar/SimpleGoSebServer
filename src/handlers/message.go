package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Message - struct
type Message struct {
	l *log.Logger
}

//NewMessage - New message
func NewMessage(l *log.Logger) *Message {
	return &Message{l}
}

func (m *Message) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	m.l.Println("Message logged...")
	data, error := ioutil.ReadAll(r.Body)

	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Something went wrong"))
		return
	}

	fmt.Fprintf(rw, "You got a message %s", data)
}

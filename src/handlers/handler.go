package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Greeting - struct
type Greeting struct {
	l *log.Logger
}

// NewGreeting - Func
func NewGreeting(l *log.Logger) *Greeting {
	return &Greeting{l}
}

// ServerHTTP - Implemented interface for Greeting type
func (g *Greeting) ServerHTTP(rw http.ResponseWriter, r *http.Request) {

	g.l.Println("Greetings logged...")
	data, error := ioutil.ReadAll(r.Body)

	if error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Something went wrong"))
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}

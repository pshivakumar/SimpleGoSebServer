package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	gh := handlers.NewGreeting(l)
	mh := handlers.NewMessage(l)

	sm := http.NewServeMux()
	sm.Handle("/", gh)
	sm.Handle("/message", mh)

	s := http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}

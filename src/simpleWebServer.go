package main

import (
	"log"
	"net/http"
	"os"
	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	gh := handlers.NewGreeting(l)

	sm := http.NewServeMux()
	sm.Handle("/", gh)

	http.ListenAndServe("127.0.0.1:9090", sm)

}

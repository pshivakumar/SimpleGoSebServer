package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		data, error := ioutil.ReadAll(r.Body)

		if error != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Something went wrong"))
			return
		}

		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.ListenAndServe("127.0.0.1:9090", nil)

}

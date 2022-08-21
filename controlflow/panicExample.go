package main

import (
	"net/http"
)

func main() {

	// Use for unrecoverable events - cannot obtain TCP port for web server
	// generates a panic when port 8080 is already in use
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}

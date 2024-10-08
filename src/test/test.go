package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		/*
			for name, values := range r.Header {
				// Loop over all values for the name.
				for _, value := range values {
					fmt.Fprintf(w, name, value)
				}
			}
		*/
		w.Header().Add("Content-Type", "application/json")
		//r.ParseForm()
		fmt.Fprintf(w, "{\"success\" : \"Welcome Home!\"}")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		/*
			for name, values := range r.Header {
				// Loop over all values for the name.
				for _, value := range values {
					fmt.Fprintf(w, name, value)
				}
			}
		*/
		w.Header().Add("Content-Type", "application/json")
		//r.ParseForm()
		fmt.Fprintf(w, "{\"success\" : \"Che Vai!\"}")
	})

	http.ListenAndServe(":9001", nil)
}

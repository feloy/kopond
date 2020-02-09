package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bonjour, vous avez demandé : %s\n", r.URL.Path)
		log.Printf("Request: %s", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}

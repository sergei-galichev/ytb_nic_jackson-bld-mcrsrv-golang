package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, World!")
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Bye, World!")
	})
	http.ListenAndServe(":9090", nil)
}

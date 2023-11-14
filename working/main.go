package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, World!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops, something went wrong", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello, %s\n", d)
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Bye, World!")
	})
	http.ListenAndServe(":9090", nil)
}

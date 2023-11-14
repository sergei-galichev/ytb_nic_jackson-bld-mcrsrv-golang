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
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Hello, %s", d)
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Bye, World!")
	})
	http.ListenAndServe(":9090", nil)
}

package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handling Hello Request")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading request body", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s\n", d)
}

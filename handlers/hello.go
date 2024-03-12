// File: handlers/hello.go

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

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		// or in 2 lines
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oops"))
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}

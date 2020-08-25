package main

import (
	"fmt"
	"net/http"
	"os"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("received request from %s", r.RemoteAddr)
	host, err := os.Hostname()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("hello world from %s", host)))
}

func main() {
	handler := myHandler{}
	server := http.Server{
		Handler: &handler,
		Addr:    ":8080",
	}
	fmt.Println("listening on :8080")
	server.ListenAndServe()
}

// docker build -t handler .
// docker run --rm -p 8080:8080 handler
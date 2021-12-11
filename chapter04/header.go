package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}
func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/headers", headers)

	server.ListenAndServe()
}

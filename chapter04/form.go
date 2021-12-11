package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form, r.PostForm)
}

func main() {

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

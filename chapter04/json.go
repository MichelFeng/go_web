package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Michel",
		Threads: []string{"first", "second"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

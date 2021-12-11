package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 负责生成HTML并写入ResponseWriter
func index(w http.ResponseWriter, r *http.Request) {

	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		}
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// 负责错误处理
func err(w http.ResponseWriter, r *http.Request) {
}

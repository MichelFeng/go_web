package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("tmpl.html")

	// rand.Seed(time.Now().Unix())
	// t.Execute(w, rand.Intn(10) > 5)

	// daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	// var daysOfWeek []string
	// t.Execute(w, daysOfWeek)

	// t.Execute(w, "hello")

	// t, _ := template.ParseFiles("t1.html", "t2.html")
	// t.Execute(w, "Hello world!")

	// funcMap := template.FuncMap{"fdate": formatDate}
	// t := template.New("tmpl.html").Funcs(funcMap)
	// t, _ = t.ParseFiles("tmpl.html")
	// t.Execute(w, time.Now())

	// t, _ := template.ParseFiles("tmpl.html")
	// content := `I asked: <i> "What's up?"</i>`
	// t.Execute(w, content)

	rand.Seed(time.Now().Unix())

	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		// t, _ = template.ParseFiles("layout.html", "blue_hello.html")
		t, _ = template.ParseFiles("layout.html")
	}

	t.ExecuteTemplate(w, "layout", "")
}

package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func main() {

	// http.ListenAndServe(":8080", nil)

	// handler := MyHandler{}
	// server := http.Server{
	// 	Addr:    "0.0.0.0:8080",
	// 	Handler: &handler,
	// }

	helloHandler := HelloHandler{}
	worldHandler := WorldHandler{}
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	// Attention：如果url不是以「/」结尾，则只会进行完全匹配，如果以「/」结尾，则会进行前缀匹配
	http.Handle("/hello", log(helloHandler))
	http.Handle("/world", log(worldHandler))

	http.HandleFunc("/hello2", log2(hello))
	http.HandleFunc("/world2", log2(world))

	server.ListenAndServe()
}

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler function called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func log2(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

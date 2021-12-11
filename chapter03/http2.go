package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandlerNew struct{}

func (h *MyHandlerNew) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandlerNew{}

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: &handler,
	}

	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("server.crt", "server.key")

	// 使用下面的命令测试
	// curl -I --http2 https://localhost:8080/  --insecure
}

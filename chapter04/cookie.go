package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/set_msg", setMessage)
	http.HandleFunc("/show_msg", showMessage)
	server.ListenAndServe()
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first",
		Value:    "Go Web",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second",
		Value:    "Go GUI",
		HttpOnly: true,
	}

	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// h := r.Header["Cookie"]
	// fmt.Fprintln(w, h)

	c1, err := r.Cookie("first")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fmt.Fprintln(w, "No message fount")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

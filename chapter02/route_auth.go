package main

import "net/http"

// 负责错误处理
func login(w http.ResponseWriter, r *http.Request) {
}

func logout(w http.ResponseWriter, r *http.Request) {
}

func signup(w http.ResponseWriter, r *http.Request) {
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		// 密码正确，跳转首页
		http.Redirect(w, r, "/", 302)
	} else {
		// 密码错误，跳转登陆页
		http.Redirect(w, r, "/login", 302)
	}
}

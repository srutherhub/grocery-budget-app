package handlers

import (
	views "app/views/pages"
	"net/http"

	v "github.com/srutherhub/web-app/views"
)

func Base() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		homepage := views.Home()
		props := v.IndexProps{Title: "Home"}
		v.Index(homepage, props).Render(r.Context(), w)
	}
}

func App() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apppage := views.App()
		props := v.IndexProps{Title: "App"}
		v.Index(apppage, props).Render(r.Context(), w)
	}
}

func Login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		loginpage := views.Login()
		props := v.IndexProps{Title: "Login"}
		v.Index(loginpage, props).Render(r.Context(), w)
	}
}

func Signup() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		signuppage := views.Signup()
		props := v.IndexProps{Title: "Signup"}
		v.Index(signuppage, props).Render(r.Context(), w)
	}
}

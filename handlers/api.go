package handlers

import (
	"app/services"
	comp "app/views/components"
	"fmt"
	"net/http"
)

func ApiSubmitChat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		msg := r.FormValue("msg")

		if msg == "" {

		}

		test := comp.IUserMessageProps{Content: msg}

		comp.UserMessage(test).Render(r.Context(), w)

	}
}

func ApiLogin(authService services.IAuthService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		token, refresh, err := authService.SignInWithEmailPassword(username, password)

		if err != nil {
			fmt.Println(err)
		}

		SetHttpCookie(w, string(token), "session")
		SetHttpCookie(w, string(refresh), "refresh")
		Redirect(w, r, "/app")
	}
}

func ApiSignup(authService services.IAuthService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		token, refresh, err := authService.SignUpWithEmailPassword(username, password)

		if err != nil {
			fmt.Println(err)
		}
		SetHttpCookie(w, string(token), "session")
		SetHttpCookie(w, string(refresh), "refresh")
		Redirect(w, r, "/app")
	}

}

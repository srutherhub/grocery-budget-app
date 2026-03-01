package handlers

import (
	comp "app/views/components"
	"net/http"
)

func SubmitChat() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		msg := r.FormValue("msg")

		if msg == "" {

		}

		test := comp.IUserMessageProps{Content: msg}

		comp.UserMessage(test).Render(r.Context(), w)

	}
}

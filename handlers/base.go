package handlers

import (
	views "app/views/pages"
	"net/http"

	v "github.com/srutherhub/web-app/views"
)

func Base(w http.ResponseWriter, r *http.Request) {
	homepage := views.Home()
	props := v.IndexProps{Title: "Home"}
	v.Index(homepage, props).Render(r.Context(), w)
}

func App(w http.ResponseWriter, r *http.Request) {
	apppage := views.App()
	props := v.IndexProps{Title: "App"}
	v.Index(apppage, props).Render(r.Context(), w)
}

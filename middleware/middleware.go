package middleware

import (
	h "app/handlers"
	s "app/services"
	"net/http"
)

func IsLoggedIn(authService s.IAuthService) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			sessionCookie, err := r.Cookie("session")
			if err != nil || sessionCookie == nil || sessionCookie.Value == "" {
				h.Redirect(w, r, "/login")
				return
			}

			if h.IsAccessTokenExpired(sessionCookie.Value) {
				h.Redirect(w, r, "/login")
				return
			}

			next(w, r)
		}
	}
}

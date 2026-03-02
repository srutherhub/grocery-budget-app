package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Redirect(w http.ResponseWriter, r *http.Request, path string) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", path)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, path, http.StatusTemporaryRedirect)
	}
}

func SetHttpCookie(w http.ResponseWriter, val string, name string) {

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    val,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})

}

func IsAccessTokenExpired(accessToken string) bool {
	token, _, err := jwt.NewParser().ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return true
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return true
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return true
	}

	return exp.Before(time.Now())

}

package controllers

import (
	"net/http"
	"web/src/cookies"
)

func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)

	http.Redirect(w, r, "/login", http.StatusFound)
}

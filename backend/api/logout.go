package handler

import (
	"net/http"

	"github.com/leoldding/cpf-wl/auth"
	"github.com/leoldding/cpf-wl/utils"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if utils.EnableCors(&w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}

	auth.Logout(w, r)
}

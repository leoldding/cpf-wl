package handler

import (
	"encoding/json"
	"net/http"

	"github.com/leoldding/cpf-wl/auth"
	"github.com/leoldding/cpf-wl/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if utils.EnableCors(&w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}

	var creds auth.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth.Login(w, r, creds)
}

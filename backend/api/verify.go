package handler

import (
	"log"
	"net/http"

	"github.com/leoldding/cpf-wl/auth"
	"github.com/leoldding/cpf-wl/utils"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	if utils.EnableCors(&w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := auth.VerifyToken(w, r); err != nil {
		log.Println(err.Error())
		return
	}
}

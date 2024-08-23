package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/leoldding/cpf-wl/auth"
	"github.com/leoldding/cpf-wl/database"
	"github.com/leoldding/cpf-wl/utils"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if utils.EnableCors(&w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := auth.VerifyToken(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	var user *database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	pool, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer pool.Close()

	err = database.UpdateUser(ctx, pool, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Success"))
}

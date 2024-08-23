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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if utils.EnableCors(&w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := auth.VerifyToken(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	var newUser *database.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	pool, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer pool.Close()

	err = database.CreateUser(ctx, pool, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

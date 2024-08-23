package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/leoldding/cpf-wl/database"
	"github.com/leoldding/cpf-wl/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if utils.EnableCors(&w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}

	ctx := context.Background()
	pool, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer pool.Close()

	users, err := database.GetUsers(ctx, pool)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

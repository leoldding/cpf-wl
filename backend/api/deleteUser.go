package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/leoldding/cpf-wl/auth"
	"github.com/leoldding/cpf-wl/database"
	"github.com/leoldding/cpf-wl/utils"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if utils.EnableCors(&w, r) {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := auth.VerifyToken(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/deleteUser/")
	log.Println(id)

	ctx := context.Background()
	pool, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer pool.Close()

	err = database.DeleteUser(ctx, pool, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Success"))
}

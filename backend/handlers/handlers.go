package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leoldding/cpf-wl/auth"
	"github.com/leoldding/cpf-wl/database"
	db "github.com/leoldding/cpf-wl/database"
)

func RegisterHandlers(router *mux.Router, ctx context.Context, pool *pgxpool.Pool) {
	log.Println("Registering Handlers")

	router.HandleFunc("/api/verify", verifyToken).Methods("GET")

	router.HandleFunc("/api/login", login()).Methods("POST")
	router.HandleFunc("/api/logout", logout).Methods("GET")

	router.HandleFunc("/api/users", checkJWT(createUser(ctx, pool))).Methods("POST")
	router.HandleFunc("/api/users", getUsers(ctx, pool)).Methods("GET")
	router.HandleFunc("/api/users", checkJWT(updateUser(ctx, pool))).Methods("PATCH")
	router.HandleFunc("/api/users/{id}", checkJWT(deleteUser(ctx, pool))).Methods("DELETE")
}

func checkJWT(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.VerifyToken(w, r); err != nil {
			log.Println(err.Error())
			return
		}

		handler(w, r)
	}
}

func verifyToken(w http.ResponseWriter, r *http.Request) {
	auth.VerifyToken(w, r)
}

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds auth.Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		auth.Login(w, r, creds)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout(w, r)
}

func createUser(ctx context.Context, pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser *db.User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := database.CreateUser(ctx, pool, newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}
}

func getUsers(ctx context.Context, pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := database.GetUsers(ctx, pool)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func updateUser(ctx context.Context, pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user *db.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := database.UpdateUser(ctx, pool, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Success"))
	}
}

func deleteUser(ctx context.Context, pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := database.DeleteUser(ctx, pool, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Success"))
	}
}

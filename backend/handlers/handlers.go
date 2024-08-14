package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/leoldding/cpf-wl/database"
)

func RegisterHandlers(router *mux.Router, database *db.Database) {
	log.Println("Registering Handlers")
	router.HandleFunc("/api/users", createUser(database)).Methods("POST")
	router.HandleFunc("/api/users", getUsers(database)).Methods("GET")
	router.HandleFunc("/api/users", updateUser(database)).Methods("PATCH")
	router.HandleFunc("/api/users/{id}", deleteUser(database)).Methods("DELETE")
}

func createUser(database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser *db.User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		database.CreateUser(newUser)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}
}

func getUsers(database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := database.GetUsers()
		json.NewEncoder(w).Encode(users)
	}
}

func updateUser(database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user *db.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		database.UpdateUser(user)
		w.Write([]byte("Success"))
	}
}

func deleteUser(database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		database.DeleteUser(id)
		w.Write([]byte("Success"))
	}
}

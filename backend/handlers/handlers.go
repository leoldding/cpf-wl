package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoldding/cpf-wl/auth"
	db "github.com/leoldding/cpf-wl/database"
)

func RegisterHandlers(router *mux.Router, database *db.Database) {
	log.Println("Registering Handlers")

	router.HandleFunc("/api/verify", verifyToken).Methods("GET")

	router.HandleFunc("/api/login", login()).Methods("POST")
	router.HandleFunc("/api/logout", logout).Methods("GET")

	router.HandleFunc("/api/users", checkJWT(createUser(database))).Methods("POST")
	router.HandleFunc("/api/users", getUsers(database)).Methods("GET")
	router.HandleFunc("/api/users", checkJWT(updateUser(database))).Methods("PATCH")
	router.HandleFunc("/api/users/{id}", checkJWT(deleteUser(database))).Methods("DELETE")

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

package main

import (
	"log"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	db "github.com/leoldding/cpf-wl/database"
	"github.com/leoldding/cpf-wl/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	database := db.NewDatabase()
	database.LoadMockData()

	handlers.RegisterHandlers(router, database)

	headersOk := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := gorillaHandlers.AllowedOrigins([]string{"http://localhost:5173"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"POST", "GET", "PATCH", "DELETE"})

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", gorillaHandlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

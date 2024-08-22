package main

import (
	"context"
	"log"
	"net/http"
	"os"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/leoldding/cpf-wl/database"
	"github.com/leoldding/cpf-wl/handlers"
)

var origins = os.Getenv("WL_ORIGIN")

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	ctx := context.Background()
	pool, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer pool.Close()

	handlers.RegisterHandlers(router, ctx, pool)

	headersOk := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := gorillaHandlers.AllowedOrigins([]string{origins})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"POST", "GET", "PATCH", "DELETE"})
	credentialsOk := gorillaHandlers.AllowCredentials()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", gorillaHandlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(router)))
}

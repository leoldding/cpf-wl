package utils

import (
	"net/http"
	"os"
)

var origin = os.Getenv("WL_ORIGIN")

func EnableCors(w *http.ResponseWriter, r *http.Request) bool {
	(*w).Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type")
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, OPTIONS")

	return r.Method == "OPTIONS"
}

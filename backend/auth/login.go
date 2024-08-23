package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var username = os.Getenv("WL_USERNAME")
var password = os.Getenv("WL_PASSWORD")
var secretKey = []byte(os.Getenv("WL_SECRET_KEY"))

func Login(w http.ResponseWriter, r *http.Request, creds Credentials) {
	if creds.Username != username || creds.Password != password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := createToken(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("SUCCESSFUL LOGIN")
	http.SetCookie(w, &http.Cookie{
		Name:     "wl-leaderboard",
		Value:    tokenString,
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
}

func createToken(user string) (string, error) {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(30 * time.Minute).Unix(),
		"user": user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("wl-leaderboard")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "JWT is missing", http.StatusUnauthorized)
			return errors.New("ERROR: MISSING JWT")
		}
		http.Error(w, "JWT is invalid", http.StatusBadRequest)
		return errors.New("ERROR: INVALID JWT")
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			http.Error(w, "User is unauthorized", http.StatusUnauthorized)
			return nil, errors.New("ERROR: UNAUTHORIZED USER / INCORRECT JWT")
		}
		return secretKey, nil

	})

	if err != nil {
		http.Error(w, "Issue parsing JWT", http.StatusInternalServerError)
		return err
	}

	if !token.Valid {
		http.Error(w, "JWT is invalid", http.StatusBadRequest)
		return errors.New("ERROR: INVALID JWT")
	}

	log.Println("TOKEN IS VERIFIED")
	return nil
}

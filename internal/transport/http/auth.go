package http

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	swt "github.com/thefueley/scholar-power-api/token"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Bearer token-string
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// fmt.Printf("authHeaderParts[1]: %v\n", authHeaderParts[1])

		if err := Verify(authHeaderParts[1]); err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else {
			original(w, r)
		}
	}
}

// Verify checks if the token is valid or not
func Verify(token string) error {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, swt.ErrInvalidToken
		}
		return []byte(os.Getenv("SCHOLAR_POWER_API_SIGNING_KEY")), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &swt.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, swt.ErrExpiredToken) {
			return swt.ErrExpiredToken
		}
		return swt.ErrInvalidToken
	}

	_, ok := jwtToken.Claims.(*swt.Payload)
	if !ok {
		return swt.ErrInvalidToken
	}

	return nil
}

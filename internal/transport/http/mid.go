package http

import (
	"context"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func JSONMiddleware(next http.Handler) http.Handler {
	if os.Getenv("API_MODE") == "dev" {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "https://deft-torte-4578e5.netlify.app")
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
			}).Info("request received")
		next.ServeHTTP(w, r)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

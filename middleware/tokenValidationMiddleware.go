package middleware

import (
	"log"
	"net/http"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		log.Println(token)
		next.ServeHTTP(w, r)
	})
}

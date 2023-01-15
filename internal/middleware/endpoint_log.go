package middleware

import (
	"log"
	"net/http"
	"time"
)

// LogCurrentRequest It's a simple middleware to log requests and time it
func LogCurrentRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[INFO] %v request to %v, took: %s", r.Method, r.URL, time.Since(start))
	})
}

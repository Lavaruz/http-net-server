package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger middleware untuk mencatat request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Catat waktu mulai
		start := time.Now()

		// Panggil handler berikutnya
		next.ServeHTTP(w, r)

		// Catat waktu selesai
		duration := time.Since(start)

		// Log request
		log.Printf(
			"Method: %s | Path: %s | Duration: %v | IP: %s",
			r.Method,
			r.URL.Path,
			duration,
			r.RemoteAddr,
		)
	})
}

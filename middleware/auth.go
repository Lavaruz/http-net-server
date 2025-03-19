package middleware

import (
	"context"
	"http-net-server/security"
	"net/http"
	"strings"
)

// Auth middleware untuk autentikasi JWT
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth untuk endpoint tertentu
		if r.URL.Path == "/login" || r.URL.Path == "/register" {
			next.ServeHTTP(w, r)
			return
		}

		// Ambil token dari header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Cek format token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Validasi token
		claims, err := security.ValidateToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Tambahkan claims ke context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// isValidToken adalah fungsi untuk validasi token
// Ini hanya contoh sederhana, Anda bisa menggantinya dengan implementasi JWT atau metode autentikasi lainnya
func isValidToken(token string) bool {
	// Implementasi validasi token
	// Contoh sederhana: token harus lebih dari 10 karakter
	return len(token) > 10
}

package middleware

import (
	"net/http"
	"strings"
)

// Auth middleware untuk autentikasi
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Cek format token (Bearer token)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		// Di sini Anda bisa menambahkan validasi token
		// Misalnya menggunakan JWT atau token lainnya
		if !isValidToken(token) {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token valid, lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}

// isValidToken adalah fungsi untuk validasi token
// Ini hanya contoh sederhana, Anda bisa menggantinya dengan implementasi JWT atau metode autentikasi lainnya
func isValidToken(token string) bool {
	// Implementasi validasi token
	// Contoh sederhana: token harus lebih dari 10 karakter
	return len(token) > 10
}

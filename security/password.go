package security

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword mengenkripsi password menggunakan bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword memverifikasi password
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatePassword memvalidasi kekuatan password
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password harus minimal 8 karakter")
	}

	// Harus mengandung huruf besar
	upper := regexp.MustCompile(`[A-Z]`)
	if !upper.MatchString(password) {
		return errors.New("password harus mengandung huruf besar")
	}

	// Harus mengandung huruf kecil
	lower := regexp.MustCompile(`[a-z]`)
	if !lower.MatchString(password) {
		return errors.New("password harus mengandung huruf kecil")
	}

	// Harus mengandung angka
	number := regexp.MustCompile(`[0-9]`)
	if !number.MatchString(password) {
		return errors.New("password harus mengandung angka")
	}

	// Harus mengandung karakter khusus
	special := regexp.MustCompile(`[!@#$%^&*]`)
	if !special.MatchString(password) {
		return errors.New("password harus mengandung karakter khusus (!@#$%^&*)")
	}

	return nil
}

// SanitizeInput membersihkan input dari karakter berbahaya
func SanitizeInput(input string) string {
	// Hapus karakter HTML
	html := regexp.MustCompile(`<[^>]*>`)
	input = html.ReplaceAllString(input, "")

	// Hapus karakter SQL injection
	sql := regexp.MustCompile(`['";]`)
	input = sql.ReplaceAllString(input, "")

	return input
}

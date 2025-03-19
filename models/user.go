package models

import (
	"database/sql"
	"fmt"
	"http-net-server/security"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"` // Password tidak akan dikirim dalam response
}

// CreateUser membuat user baru dengan password yang di-hash
func CreateUser(db *sql.DB, user *User) error {
	// Validasi password
	if err := security.ValidatePassword(user.Password); err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Sanitize input
	user.Name = security.SanitizeInput(user.Name)
	user.Email = security.SanitizeInput(user.Email)

	// Query dengan prepared statement untuk mencegah SQL injection
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	result, err := db.Exec(query, user.Name, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
	fmt.Println(user)
	if err != nil {
		return err
	}

	user.ID = id
	user.Password = "" // Hapus password dari struct setelah disimpan
	return nil
}

// GetUser mengambil user berdasarkan ID
func GetUser(db *sql.DB, id int64) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email FROM users WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers mengambil semua user
func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(rows)

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// AuthenticateUser memverifikasi kredensial user
func AuthenticateUser(db *sql.DB, email, password string) (*User, error) {
	var user User
	var hashedPassword string

	query := `SELECT id, name, email, password FROM users WHERE email = ?`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &hashedPassword)
	if err != nil {
		return nil, err
	}

	if !security.CheckPassword(password, hashedPassword) {
		return nil, sql.ErrNoRows
	}

	user.Password = "" // Hapus password dari struct
	return &user, nil
}

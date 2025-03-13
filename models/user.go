package models

import (
	"database/sql"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser membuat user baru di database
func CreateUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
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

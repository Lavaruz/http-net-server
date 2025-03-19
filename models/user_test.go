package models

// import (
// 	"database/sql"
// 	"testing"
// )

// // MockDB adalah mock database untuk testing
// type MockDB struct {
// 	users []User
// }

// func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
// 	// Implementasi mock query
// 	return nil, nil
// }

// func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
// 	// Implementasi mock query row
// 	return nil
// }

// func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
// 	// Implementasi mock exec
// 	return nil, nil
// }

// func TestCreateUser(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		user    User
// 		wantErr bool
// 	}{
// 		{
// 			name: "Valid user",
// 			user: User{
// 				Name:  "Test User",
// 				Email: "test@example.com",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Empty name",
// 			user: User{
// 				Email: "test@example.com",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Empty email",
// 			user: User{
// 				Name: "Test User",
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockDB := &MockDB{}
// 			err := CreateUser(mockDB, &tt.user)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		id      int64
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Valid ID",
// 			id:      1,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "Invalid ID",
// 			id:      -1,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockDB := &MockDB{}
// 			_, err := GetUser(mockDB, tt.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestGetAllUsers(t *testing.T) {
// 	mockDB := &MockDB{}
// 	users, err := GetAllUsers(mockDB)
// 	if err != nil {
// 		t.Errorf("GetAllUsers() error = %v", err)
// 	}
// 	if users == nil {
// 		t.Error("GetAllUsers() returned nil users")
// 	}
// }

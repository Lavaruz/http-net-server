package integration

import (
	"bytes"
	"encoding/json"
	"http-net-server/config"
	"http-net-server/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestDB(t *testing.T) {
	// Inisialisasi test database
	config.InitDB()
}

func TestUserAPI(t *testing.T) {
	setupTestDB(t)

	t.Run("Create User", func(t *testing.T) {
		user := models.User{
			Name:  "Test User",
			Email: "test@example.com",
		}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Buat mux router dan handler
		mux := http.NewServeMux()
		mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var user models.User
				if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				if err := models.CreateUser(config.DB, &user); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(user)
			}
		})

		mux.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status %v; got %v", http.StatusCreated, w.Code)
		}
	})

	t.Run("Get All Users", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				users, err := models.GetAllUsers(config.DB)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				json.NewEncoder(w).Encode(users)
			}
		})

		mux.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %v; got %v", http.StatusOK, w.Code)
		}
	})
}

func TestErrorHandling(t *testing.T) {
	setupTestDB(t)

	t.Run("Invalid JSON", func(t *testing.T) {
		body := bytes.NewBufferString("{invalid json}")
		req := httptest.NewRequest("POST", "/users", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var user models.User
				if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
		})

		mux.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %v; got %v", http.StatusBadRequest, w.Code)
		}
	})
}

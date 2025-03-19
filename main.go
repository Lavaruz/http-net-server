package main

import (
	"encoding/json"
	"http-net-server/config"
	"http-net-server/errors"
	"http-net-server/middleware"
	"http-net-server/models"
	"log"
	"net/http"
	"strconv"
)

// Struct untuk response
type Response struct {
	Message string      `json:"message"`
	Method  string      `json:"method"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	// Inisialisasi database
	config.InitDB()

	// Buat mux router
	mux := http.NewServeMux()

	// Handler untuk endpoint root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r)
		case http.MethodPost:
			handlePost(w, r)
		case http.MethodPut:
			handlePut(w, r)
		case http.MethodDelete:
			handleDelete(w, r)
		default:
			errors.WriteError(w, errors.ErrBadRequest)
		}
	})

	// Handler untuk users (dengan auth)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Ambil semua users
			users, err := models.GetAllUsers(config.DB)
			if err != nil {
				errors.WriteError(w, errors.NewError(http.StatusInternalServerError, "Failed to get users", err))
				return
			}

			response := Response{
				Message: "Daftar semua users",
				Method:  "GET",
				Status:  http.StatusOK,
				Data:    users,
			}
			sendJSON(w, response)

		case http.MethodPost:
			var user models.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				errors.WriteError(w, errors.NewError(http.StatusBadRequest, "Invalid request body", err))
				return
			}

			// Validasi input
			if user.Name == "" || user.Email == "" {
				errors.WriteError(w, errors.NewError(http.StatusUnprocessableEntity, "Name and email are required", nil))
				return
			}

			log.Println(user)

			if err := models.CreateUser(config.DB, &user); err != nil {
				errors.WriteError(w, errors.NewError(http.StatusInternalServerError, "Failed to create user", err))
				return
			}

			response := Response{
				Message: "User berhasil dibuat",
				Method:  "POST",
				Status:  http.StatusCreated,
				Data:    user,
			}
			sendJSON(w, response)
		}
	})

	// Handler untuk user detail (dengan auth)
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// Ambil ID dari URL
		idStr := r.URL.Path[len("/users/"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			errors.WriteError(w, errors.NewError(http.StatusBadRequest, "Invalid user ID", err))
			return
		}

		if r.Method == http.MethodGet {
			user, err := models.GetUser(config.DB, id)
			if err != nil {
				errors.WriteError(w, errors.ErrNotFound)
				return
			}

			response := Response{
				Message: "Detail user",
				Method:  "GET",
				Status:  http.StatusOK,
				Data:    user,
			}
			sendJSON(w, response)
		}
	})

	// Buat handler dengan middleware
	// handler := middleware.Logger(middleware.CORS(middleware.Auth(mux)))
	handler := middleware.Logger(middleware.CORS(mux))

	// Menjalankan server di port 8080
	log.Println("Server berjalan di http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

// Fungsi helper untuk mengirim JSON
func sendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Ini adalah response GET",
		Method:  "GET",
	}
	sendJSON(w, response)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		errors.WriteError(w, errors.NewError(http.StatusBadRequest, "Failed to parse form", err))
		return
	}

	// Ambil data dari form
	name := r.FormValue("name")
	email := r.FormValue("email")

	// Validasi input
	if name == "" || email == "" {
		errors.WriteError(w, errors.NewError(http.StatusUnprocessableEntity, "Name and email are required", nil))
		return
	}

	response := Response{
		Message: "Data berhasil diterima",
		Method:  "POST",
		Data: map[string]string{
			"name":  name,
			"email": email,
		},
	}
	sendJSON(w, response)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Ini adalah response PUT",
		Method:  "PUT",
	}
	sendJSON(w, response)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Ini adalah response DELETE",
		Method:  "DELETE",
	}
	sendJSON(w, response)
}

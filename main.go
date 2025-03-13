package main

import (
	"encoding/json"
	"http-net-server/config"
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

	// Handler untuk endpoint root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Ambil semua users
			users, err := models.GetAllUsers(config.DB)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = models.CreateUser(config.DB, &user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// Ambil ID dari URL
		idStr := r.URL.Path[len("/users/"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "ID tidak valid", http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodGet {
			user, err := models.GetUser(config.DB, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
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

	// Menjalankan server di port 8080
	log.Println("Server berjalan di http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
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
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Ambil data dari form
	name := r.FormValue("name")
	email := r.FormValue("email")

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

package middleware

import (
	"encoding/json"
	"http-net-server/config"
	"net/http"
	"time"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// Check database connection
	if err := config.DB.Ping(); err != nil {
		http.Error(w, "Database unhealthy", http.StatusServiceUnavailable)
		return
	}

	response := map[string]string{
		"status": "healthy",
		"time":   time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

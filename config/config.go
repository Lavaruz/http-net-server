package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Config menyimpan semua konfigurasi aplikasi
type Config struct {
	// Database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	// Server
	ServerPort int
	ServerHost string

	// Environment
	Environment string

	// JWT
	JWTSecret string
	JWTExpiry int // dalam jam

	// Database connection
	DB *sql.DB
}

// InitDB menginisialisasi koneksi database
func InitDB() {
	config, err := LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Buat connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	// Buka koneksi database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Test koneksi
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	DB = db
}

// LoadConfig memuat konfigurasi dari environment variables
func LoadConfig() (*Config, error) {
	config := &Config{}

	// Load environment
	config.Environment = getEnvOrDefault("APP_ENV", "development")

	// Database config
	config.DBHost = getEnvOrDefault("DB_HOST", "localhost")
	config.DBPort = getEnvAsIntOrDefault("DB_PORT", 3306)
	config.DBUser = getEnvOrDefault("DB_USER", "root")
	config.DBPassword = getEnvOrDefault("DB_PASSWORD", "181001")
	config.DBName = getEnvOrDefault("DB_NAME", "go_api")

	// Server config
	config.ServerPort = getEnvAsIntOrDefault("SERVER_PORT", 8080)
	config.ServerHost = getEnvOrDefault("SERVER_HOST", "localhost")

	// JWT config
	config.JWTSecret = getEnvOrDefault("JWT_SECRET", "your-secret-key")
	config.JWTExpiry = getEnvAsIntOrDefault("JWT_EXPIRY", 24)

	return config, nil
}

// Helper functions
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

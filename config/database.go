package config

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// // InitDB menginisialisasi koneksi database
// func InitDB() {
// 	config, err := LoadConfig()
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to load config: %v", err))
// 	}

// 	// Buat connection string
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
// 		config.DBUser,
// 		config.DBPassword,
// 		config.DBHost,
// 		config.DBPort,
// 		config.DBName,
// 	)

// 	// Buka koneksi database
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to connect to database: %v", err))
// 	}

// 	// Test koneksi
// 	if err := db.Ping(); err != nil {
// 		panic(fmt.Sprintf("Failed to ping database: %v", err))
// 	}

// 	config.DB = db
// }

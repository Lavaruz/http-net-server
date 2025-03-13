package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// Konfigurasi koneksi database
	dsn := "root:181001@tcp(localhost:3306)/go_api?parseTime=true"

	// Membuka koneksi ke database
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error koneksi ke database:", err)
	}

	// Test koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error ping database:", err)
	}

	fmt.Println("Berhasil terhubung ke database!")
}

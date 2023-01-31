package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := newDB()
	defer db.Close()

	var n int
	if err := db.QueryRow(`select 1 + 1`).Scan(&n); err != nil {
		panic(err)
	}
	log.Println("n:", n)
}

func newDB() *sql.DB {
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("failed to init mysql:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping:", err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

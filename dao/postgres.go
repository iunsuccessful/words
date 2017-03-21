package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
    "log"
)

func getDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:635307098@localhost/word?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

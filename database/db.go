package database

import (
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var DB *sqlx.DB

// InitDB initializes the PostgreSQL connection
func InitDB() {
    var err error
    connStr := "user=postgres password=postgres dbname=go_crud_data sslmode=disable"
    DB, err = sqlx.Connect("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }
}

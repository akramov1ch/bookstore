package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/akramov1ch/bookstore/internal/config"
	_ "github.com/lib/pq"
)


func DbConnect() *sql.DB {
	con, err := config.LoadConfig()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        con.DBHost,
		con.DBPort,
		con.DBUser,
		con.DBPassword,
		con.DBName,
    )
    DB, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error connecting to the database: %s", err)
    }
	return DB
}

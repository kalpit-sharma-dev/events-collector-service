package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	HOST = "database"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func GetDatabaseProvider(username, password, database string) *sql.DB {
	//db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return conn
	}
	//db.Conn = conn
	err = conn.Ping()
	if err != nil {
		return conn
	}
	log.Println("Database connection established")
	return conn
}

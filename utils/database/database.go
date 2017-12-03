package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect(user, password, dbname, host, port string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
		user, password, dbname, host, port)
	return sql.Open("postgres", connStr)
}

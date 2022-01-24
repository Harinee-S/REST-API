package modules

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "0712"
	DB_NAME     = "Signup"
)

// DB set up
func setupDB() *sql.DB {
	var db *sql.DB

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

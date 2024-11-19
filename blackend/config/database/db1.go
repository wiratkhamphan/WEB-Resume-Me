package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect1() (*sql.DB, error) { // Return *sql.DB
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "laravel_dome"
	db, err := sqlx.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		return nil, err
	}

	// Ensure the database connection is valid
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db.DB, nil // Return the underlying *sql.DB
}

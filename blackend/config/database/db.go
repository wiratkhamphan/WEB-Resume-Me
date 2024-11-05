package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() (*sqlx.DB, error) {
	// godotenv.Load()
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "shoplek"
	db, err := sqlx.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		return nil, err
	}

	// Ensure the database connection is valid
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil

}

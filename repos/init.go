package repos

import (
	"fmt"

	"database/sql"
)

var db *sql.DB

const (
	DB_NAME     = "nhaoday_development"
	DB_USER     = "tybui"
	DB_PASSWORD = "p@$$w0rd"
)

func init() {
	var err error
	connection_info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	if db == nil {
		db, err = sql.Open("postgres", connection_info)
		if err != nil {
			panic(err)
		}

		if err = db.Ping(); err != nil {
			panic(err)
		}
	}
}

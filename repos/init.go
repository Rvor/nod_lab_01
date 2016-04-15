package repos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

type AppConfig struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	config := new(AppConfig)
	config_file := os.Getenv("NOD_CONFIG_PATH")

	config_file = config_file + "config.json"

	f, e := os.Open(config_file)
	if e != nil {
		log.Fatal(e)
	}

	e = json.NewDecoder(f).Decode(config)
	if e != nil {
		log.Fatal(e)
	}

	connection_info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.Username, config.Password, config.Database)

	if db == nil {
		db, e = sql.Open("postgres", connection_info)
		if e != nil {
			log.Fatal(e)
		}

		if e = db.Ping(); e != nil {
			log.Fatal(e)
		}
	}
}

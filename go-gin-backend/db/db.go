package db

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectToDB() {

	envFile, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("mysql", envFile["CONNECTION_URI"])

	if err != nil {
		log.Printf("ERROR %s when opening DB\n", err)
		return
	}

}

func CloseDB(db *sql.DB) {
	db.Close()
}

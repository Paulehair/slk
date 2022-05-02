package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func init() {
	loadDBClient()
}

func loadDBClient() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	for {
		log.Println("Connecting to DB...")
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Println(err.Error())
			log.Println("Connection failled, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}
		DBClient = db
		break
	}

	log.Println("Connexion to DB successful")

	return nil
}

func GetDBClient() *sql.DB {
	if DBClient.Ping() != nil {
		loadDBClient()
		return DBClient
	} else {
		return DBClient
	}
}

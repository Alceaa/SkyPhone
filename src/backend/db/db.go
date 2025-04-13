package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func open() *sqlx.DB {
	//err := gotenv.Load()
	//if err != nil {
	//	log.Fatalf("Failed to load env. Err: %s", err)
	//}

	//dbUsername := os.Gotenv("DATABASE_USERNAME")
	//dbName := os.Gotenv("DATABASE_NAME")
	//dbPassword := os.Gotenv("DATABASE_PASSWORD")
	dbUsername := "postgres"
	dbName := "skyphone"
	dbPassword := "postgres"
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=localhost sslmode=disable", dbUsername, dbName, dbPassword)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var DB *sqlx.DB = open()

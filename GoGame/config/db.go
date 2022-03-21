package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Dbconnect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/game")

	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	return db
}

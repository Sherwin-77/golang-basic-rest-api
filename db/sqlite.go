package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var maxConns = 5
var conns = make(chan *sql.DB, maxConns)

func InitDB(pathName string) {
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		file, err := os.Create(pathName)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	for i := 0; i < maxConns; i++ {
		db, err := sql.Open("sqlite3", pathName)
		if err != nil {
			log.Fatal(err)
		}

		conns <- db
	}
}

func GetDB() *sql.DB {
	return <-conns
}

func ReleaseDB(db *sql.DB) {
	conns <- db
}

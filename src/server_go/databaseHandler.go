package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "os"
import "fmt"
import "log"

func initdb() {
	db, err := sql.Open("sqlite3","./user.db")
	if err != nil {
		fmt.Println("Database cannot be opened, exit...")
		log.Fatal(err)
		os.Exit(-1)
	}

	defer db.Close()

	checkTableExists := "SELECT name FROM sqlite_master WHERE type='table' AND name='user';"
	res, err := db.Exec(checkTableExists)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

    fmt.Println(res)
}

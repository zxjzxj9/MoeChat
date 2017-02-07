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
	res := db.QueryRow(checkTableExists)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

    var value interface{}
    err := res.Scan(&value)

    switch err {
        case sql.ErrNoRows:
            fmt.Println("No Table user detected, start building table")
        case nil:
            fmt.Println("Building database requires dropping user table, continue (y/N) ?")
            rep := "N"
            fmt.Scanln(rep)
            switch rep {
                case 'y':
                    dropTable := "DROP TABLE user;"
                    res, err := db.Exec(dropTable)
                    if err != nil {
                        fmt.Println("Dropping table error, exiting...")
                        os.Exit(-1)
                    }
                case 'N':
                    os.Exit(0)
            }
        default:
            fmt.Println("Unkown error, exiting")
            os.Exit(-1)
    }

    res, err :=  db.Exec(" CREATE TABLE user  "
                         "( uid INTEGER,      "
                         "  passwd INTEGER,   "
                         "  login_time TEXT,  "
                         "  online INTEGER,   "
                         "  last_check TEXT );"
                         )
}

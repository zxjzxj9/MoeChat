package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "os"
import "bufio"
import "fmt"
import "log"

// For simplicity, the database backends is sqlite3
// Other backends can be included with modification

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
    err = res.Scan(&value)

    switch err {
        case sql.ErrNoRows:
            fmt.Println("No Table user detected, start building table")
        case nil:
            fmt.Print("Building database requires dropping user table, continue (y/N)? ")
            input := bufio.NewScanner(os.Stdin)
            input.Scan()
            switch input.Text() {
                case "y":
                    dropTable := "DROP TABLE user;"
                    _, err := db.Exec(dropTable)
                    if err != nil {
                        fmt.Println("Dropping table error, exiting...")
                        os.Exit(-1)
                    }
                case "N":
                    os.Exit(0)
                default:
                    fmt.Println("Invalid input, exiting...")
                    os.Exit(-1)
            }
        default:
            fmt.Println("Unkown error, exiting")
            os.Exit(-1)
    }

    _, err =  db.Exec(" CREATE TABLE user  "  +
                         "( uid INTEGER,      " +
                         "  passwd INTEGER,   " +
                         "  login_time TEXT,  " +
                         "  online INTEGER,   " +
                         "  last_check TEXT);" )

    if err != nil {
        fmt.Println("Creat table failure!")
    }

    fmt.Println("Successfully create table user")
}

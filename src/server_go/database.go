package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "os"
import "time"
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

    _, err =  db.Exec(  " CREATE TABLE user  "  +
                         "( uid INTEGER,      " + // user id, primary key
                         "  uname TEXT,       " + // username
                         "  passwd INTEGER,   " + // password
                         "  login_time TEXT,  " + // when is user login?
                         "  online INTEGER,   " + // 0: offline 1: online
                         "  sid TEXT,         " + // session key (id)
                         "  auxport INTEGER,  " + // auxalliary port
                         "  last_check TEXT);" )  // last heartbeat time

    if err != nil {
        fmt.Println("Creat table failure!")
    }

    fmt.Println("Successfully create table user")
}

// Validate the user and passwd
func validate(user, passwd string) bool {
    var uid int
	db, err := sql.Open("sqlite3","./user.db")
	err = db.QueryRow(" SELECT uid FROM user " +
			          " WHERE uname = ?      " +
                      " AND   passwd = ? ;", user, passwd).Scan(&uid)

	defer db.Close()

	switch {
		case err == sql.ErrNoRows:
			log.Printf("Incorrect username or password.\n")
			return false
		case err != nil:
			log.Fatal(err)
			return false
		default:
			log.Printf("Login success!\n")
			return true
	}

}

// Update user session id for validation
func activate(user, session string) error {
    datetime := time.Now().UTC()
	db, err := sql.Open("sqlite3","./user.db")
    _, err = db.Exec(" UPDATE user SET sid = ?, login_time = ?, online = 1, last_check = ? WHERE uname = ?;",
                  session, datetime.Format(time.RFC3339), datetime.Format(time.RFC3339), user)
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }
    return err
}

// Check whether the client is still online
func checkAlive(user, sessionId string) bool {
    datetime := time.Now().UTC()
	db, err := sql.Open("sqlite3","./user.db")
	var s time.Time
    err = db.QueryRow(" SELECT last_check FROM user " +
                      " WHERE uname = ? "      +
                      " AND   sid = ?  ; ", user, sessionId).Scan(&s)
    defer db.Close()
    if err != nil {
        log.Fatal(err)
        return false
    }

	// If no reply for more than 10 mins, treat it offline
	if (datetime.Sub(s))*time.Minute > 10 {
		return false
	}

    return true
}

func getUsers() []string {
	db, err := sql.Open("sqlite3","./user.db")
	rows, err := db.Query(" SELECT uname FROM user; ")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    ret := make([]string,0)
    for rows.Next() {
        var u string
        if err := rows.Scan(&u); err != nil {
            log.Fatal(err)
        }
        ret = append(ret, u)
    }

    if err = rows.Err(); err != nil {
            log.Fatal(err)
    }
    return ret
}

func killOnline(uname, sessionId string) error {

	db, err := sql.Open("sqlite3","./user.db")
	_, err = db.Exec(" UPDATE user SET sid = null WHERE uname = ? AND sid = ?", uname, sessionId)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return err
}

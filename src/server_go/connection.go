package main

import "net"
import "fmt"
import "log"
//import "bytes"
import "strconv"
import "errors"
import "math/rand"
import "io/ioutil"
import "encoding/json"


const (
    MAX_BUFF_LEN = 1024
	SECONDARY_PORT = 7346
    letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sessLen = 24 // 24 digits session length
    heartbeat = 600 // set heartbeat time, default 600
)

// Define the message format
type message struct {
    src string
    dest string
    err error
}

// Mainloop tfor tcp connection
func runServer(addr string, port int) error {
	lstn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}
	defer lstn.Close()

    m := make(chan message)

	for {
		conn, err :=  lstn.Accept()
        log.Println("Incoming connection from: ", conn.RemoteAddr())
		if err != nil {
			return err
		}
        go comm(conn, m)
		defer conn.Close()
	}

    return nil
}

// This function run the port to send message from server to client
// Client is responsible to maintain this link
func runClient(addr string) error {
	return nil
}

// communicate with the client, main logic
func comm(conn net.Conn, msgQueue chan message) {
    // First send connection message
    // init the connection
    for {
	    recvbyte, err := ioutil.ReadAll(conn)
	    if err != nil {
		    log.Fatal(err)
		    return
	    }

	    defer conn.Close()

	    var data map[string]interface{}
	    err = json.Unmarshal(recvbyte, data)
	    if err != nil {
		    log.Fatal(err)
		    return
	    }
        // Process the data according to the protocal
	    switch data["status"].(string) {
		    case "q":
			    // do query hadling, login the client
				msg := data["info"].(map[string]string)

                switch msg["req"] {
                    case "login":
                        if login(msg, conn) != nil {
                            // returning messages
                            m := make(map[string]interface{})
							m["status"] = "r"
							m["status_code"] = 20
                            mdata := make(map[string]string)
							mdata["error"] = "Invalid username or password"
							m["info"] = mdata
							reply, err := json.Marshal(m)
							_, err = conn.Write(reply)
							if err != nil {
								log.Fatal(err)
							}
							return
                        }
					// Query online users, should firstly verify sessions
                    // Following all operations should carry a session id
					case "users":
						if checkLogin(msg) {
							ulist := getUsers()
							m := make(map[string] interface{})
							m["status"] = "r"
							m["status_code"] = 30
							mdata := make(map[string] interface{})
							mdata["ulist"] = ulist
							m["info"] = mdata
							reply, err := json.Marshal(m)
							_, err = conn.Write(reply)
							if err != nil {
								log.Fatal(err)
							}
						} else {
							m := make(map[string] interface{})
							m["status"] = "e"
							m["status_code"] = 40
							mdata := make(map[string]string)
							mdata["error"] = "Invalid username or sessionid"
							m["info"] = mdata
							reply, err := json.Marshal(m)
							_, err = conn.Write(reply)
							if err != nil {
								log.Fatal(err)
							}
							return
						}
					case "logout":
						if checkLogin(msg) {
							if err := logout(msg, conn); err != nil {
								log.Fatal(err)
							}
							return

						} else {

						}

					default:
						return

                }

            case "m":
			    // sending messages
                // firstly we will check whether arrivable
            default:
			    // return error, and exit
				return
	     }
    }
}

func comm2(conn net.Conn, m chan message) {

}

func cleaner(m chan message) {
    // Periodically check un arrivable message and sweep them out in the queue
    for {

    }
}

//@deprecated, each function should log error itself
//func logger(m chan message) {
    // Handle the error during server run
//    log.Println("Init server logger")
//}
// Function for generate session key

func randSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func login(msg map[string] string, conn net.Conn) error {
    user := msg["user"]
    passwd := msg["passwd"]

    // Call the database function to validater user passwd
	if checkLogin(msg) {
		log.Fatal("User is already online")
		return errors.New("User is already online...")
	}

    if validate(user, passwd) {
        // Sending the secondary port of server
        m := make(map[string]interface{})
        m["status"] = "r"
        m["status_code"] = 30
        infomap := make(map[string]string)
        m["info"] = infomap
        //Listen to secondary, for ending informations
        infomap["port"] = strconv.Itoa(SECONDARY_PORT)
        infomap["session"] = randSeq(sessLen)
        // Marshal the map
        reply, err := json.Marshal(m)

        // Register and activate the user session
        if err = activate(user, infomap["session"]); err != nil {
            log.Fatal(err)
            return err
        }
        if err != nil {
            log.Fatal(err)
            return err
        }
        _, err = conn.Write(reply)
        if err != nil {
            log.Fatal(err)
            return err
        }

        return nil
    } else {
        ret := errors.New("Cannot validate user!")
        log.Fatal(ret)
        return ret
    }
}

func logout(msg map[string] string, conn net.Conn ) error {
	uname, hasUserName := msg["uname"];
	sessionId, hasSessionId := msg["sessionId"];

	if hasSessionId && hasUserName {
		return killOnline(uname, sessionId)
	} else {
		return errors.New("No such user!")
	}
}


func checkLogin(msg map[string] string) bool {

    uname, hasUserName := msg["uname"];
	sessionId, hasSessionId := msg["sessionId"];

	if hasSessionId && hasUserName {
		return checkAlive(uname, sessionId)
	} else {
		return false
	}
}


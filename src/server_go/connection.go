package main

import "net"
import "fmt"
import "log"
import "bytes"
import "ioutil"
import "encoding/json"

const (
    MAX_BUFF_LEN = 1024
	SECONDARY_PORT = 7346
    letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	sessLen = 24 // 24 digits session length
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
		if err != nil {
			return err
		}
        go comm(conn, m)
		defer conn.Close()
	}

    return nil
}

// communicate with the client, main logic
func comm(conn net.Conn, m chan message) {
    // First send connection message
    // init the connection
	recvbyte, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	var data map[string]
	err = Unmarshal(recvbyte, data)
	if err != nil {
		log.Fatal(err)
		return
	}
    // Process the data according to the protocal
	switch data["status"] {
		case "q":
			// do query hadling, login the client
			msg := data.(map[string])
			user := msg["user"]
			passwd := msg["passwd"]
			// Call the database function to validater user passwd
			if validate(user, passwd) {
				// Sending the secondary port of server
				m := make(map[string], interface{})
				m["status"] = "r"
				infomap := make(map[string], string)
				m["info"] = infomap
				//Listen to secondary, for ending informations
				infomap["port"] := SECONDARY_PORT
                infomap["session"] := randSeq(sessLen)
				// UnMarshal the map

			} else {
				// Fail to login
			}
		case "m":
			// sending messages
		default:
			// return error, and exit

	}

}

func logger(m chan message) {
    // Handle the error during server run
    log.Println("Init server logger")
}

// Function for generate session key
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

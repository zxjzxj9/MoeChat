package main

import "net"
import "fmt"
import "log"
import "bytes"
import "ioutil"
import "encoding/json"

const (
    MAX_BUFF_LEN = 1024
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

}

func logger(m chan message) {
    // Handle the error during server run
    log.Println("Init server logger")
}

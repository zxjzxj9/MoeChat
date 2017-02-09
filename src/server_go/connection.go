package main

import "net"
import "fmt"

// Define the message format
type message {
    src string,
    dest string,
    err error
}

// Mainloop tfor tcp connection
func runServer(addr string, port int) err error {
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
}

// communicate with the client, main logic
func comm(conn Conn, m chan message) {
    // First send connection message
}


package main

import "net"
import "fmt"


// Mainloop tfor tcp connection
func runServer(addr string, port int) err error {
	lstn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}
	defer lstn.Close()

	for {
		conn, err :=  lstn.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()

		err = comm(conn, addr)
	}
}

// communicate with the client, main logic
func comm(conn Conn) {

}


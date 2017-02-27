package main

import "os"

//import "sys"
import "log"
import "flag"
import "fmt"

//import "strconv"

// Entry point to start the server
func main() {

	var addr = flag.String("ip", "0.0.0.0", "Moecat Server Binding Ip Address")
	var port = flag.Int("port", 3541, "Moecat Server Binding Port")
	var lserv = flag.Bool("run", false, "Start the server")
	var ldbinit = flag.Bool("initdb", false, "Initialize the database")

	flag.Parse()

	if *ldbinit {
		fmt.Println("Initializing database...")
		initdb()
		os.Exit(0)
	}

	if *lserv {
		fmt.Printf("Server will run on %s, port %d\n", *addr, *port)
		err := runServer(*addr, *port)
		if err != nil {
			log.Fatal("Server error! " + err.Error())
			os.Exit(-1)
		}
		os.Exit(0)
	}

	fmt.Println("Please use -h/--help for more information")
}

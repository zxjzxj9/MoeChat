package main

import "os"
//import "sys"
import "flag"
import "fmt"
//import "strconv"

// Entry point to start the server
func main() {

	//if len(os.Args) != 3 {
	//	fmt.Printf("Usage: ./moeserver addr port \n")
	//	os.Exit(-1)
	//}

	//addr := os.Args[1]
	//port, err := strconv.Atoi(os.Args[2])

	//if err != nil {
	//	fmt.Printf("Invalid Port %s!\n",os.Args[2])
	//	os.Exit(-1)
	//}

	//fmt.Printf("Test %d", add(1,2))
	//fmt.Printf("Server will run on %s, port %d\n",addr,port)

	var addr = flag.String("ip", "0.0.0.0", "Moecat Server Binding Ip Address")
	var port = flag.Int("port", 3541, "Moecat Server Binding Port")
	var lserv = flag.Bool("run", false, "Start the server")
	var ldbinit = flag.Bool("dbinit", false, "Initialize the database")

	flag.Parse()

	if ldbinit {
		fmt.Printf("Initializing database...")
		os.Exit(0)
	}

	if lserv {
		fmt.Printf("Server will run on %s, port %d\n", *addr, *port)
		os.Exit(0)
	}


}

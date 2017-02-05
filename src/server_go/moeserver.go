package main

import "os"
//import "sys"
import "fmt"
import "strconv"

// Entry point to start the server
func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Usage: ./moeserver addr port \n")
		os.Exit(-1)
	}

	addr := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Printf("Invalid Port %s!\n",os.Args[2])
		os.Exit(-1)
	}

	fmt.Printf("Server will run on %s, port %d\n",addr,port)

}

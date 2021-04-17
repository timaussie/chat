package main

import (
	"log"
	"net"
	"os"
)

func main() {
	log.Println("Launching server..")

	_, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// create an instance of server's core
	server_hub := newHub()
	// start accepting client
}

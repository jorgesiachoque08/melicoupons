package main

import (
	"os"

	"github.com/jorgesiachoque08/melicoupons/server"
)

const defaultPort = "8080"

func main() {

	serverPort := os.Getenv("PORT")
	//validate that the PORT environment variable is set, otherwise port 8080 is taken by default.
	if serverPort == "" {
		serverPort = defaultPort
	}

	s := server.Server{}
	//the server is initialized
	s.Initialize()
	//the server is running
	s.Run(serverPort)
}

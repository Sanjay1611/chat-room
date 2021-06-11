package main

import (
	"log"

	"github.com/sanjay/roam/pkg/server"
)

// Constant to define the port number
const (
	Port = ":3000"
)

func main() {
	s, err := server.InitializeServer()
	if err != nil {
		log.Println("Cannnot start server error: ", err)
	}
	s.Run(Port)
}

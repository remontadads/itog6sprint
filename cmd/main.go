package main

import (
	"log"
	"os"

	"go1fl-sprint6-final/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE: ", log.LstdFlags)

	srv := server.New(logger)

	logger.Println("Starting server on :8080")
	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}
}

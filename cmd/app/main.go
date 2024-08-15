package main

import (
	"log"

	"github.com/Evilcmd/Hackup-backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	server := server.NewServer()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error starting the server")
	}
}

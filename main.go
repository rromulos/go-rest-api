package main

import (
	"github.com/rromulos/go-rest-api/server"
	"github.com/rromulos/go-rest-api/database"
)

func main() {
	database.StartDatabase()
	server := server.NewServer()
	server.Run()
}

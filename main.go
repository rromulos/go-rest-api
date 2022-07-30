package main

import "github.com/rromulos/go-rest-api/server"

func main() {
	database.StartDatabase()
	server := server.NewServer()
	server.Run()
}

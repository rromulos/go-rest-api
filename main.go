package main

import "github.com/rromulos/go-rest-api/server"

func main() {
	server := server.NewServer()
	server.Run()
}

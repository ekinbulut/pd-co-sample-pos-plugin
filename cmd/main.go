package main

import (
	"log"
	"pos-plugin/cmd/server"
)

func main() {
	server := server.NewServer()
	log.Fatal(server.Start())
}

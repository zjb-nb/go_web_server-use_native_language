package main

import (
	"afterclass/server"
	"afterclass/server/router"
	"log"
)

func main() {
	server := server.NewServer("Web")
	// server.Router("GET","/", router.Home)
	server.Router("GET", "/home", router.Home)
	server.Router("GET", "/sign", router.Sign)
	log.Fatal(server.Start(":8080"))
}

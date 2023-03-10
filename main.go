package main

import (
	"afterclass/server"
	"afterclass/server/router"
	"log"
)

func main() {
	server := server.NewServer("Web")
	server.Router("/", router.Home)
	server.Router("/sign", router.Sign)
	log.Fatal(server.Start(":8080"))
}

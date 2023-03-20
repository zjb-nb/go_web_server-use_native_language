package main

import (
	"afterclass/server"
	"afterclass/server/router"
	"fmt"
	"log"
)

func main() {
	defer func() {
		if data := recover(); data != nil {
			log.Fatal(data)
		}
		fmt.Println("Program is ending......")
	}()
	servers := server.NewServer("Web", server.ComputeTimeBuilder, server.SayByeBuilder)
	// server.Router("GET","/", router.Home)
	go func() {
		server.WaitShutDone()
	}()

	servers.Router("GET", "/home", router.Home)
	servers.Router("GET", "/sign", router.Sign)
	err := servers.Start(":8080")
	if err != nil {
		panic(err)
	}
}

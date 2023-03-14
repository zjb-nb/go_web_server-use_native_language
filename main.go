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
	server := server.NewServer("Web", server.ComputeTimeBuilder, server.SayByeBuilder)
	// server.Router("GET","/", router.Home)
	server.Router("GET", "/home", router.Home)
	server.Router("GET", "/sign", router.Sign)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}

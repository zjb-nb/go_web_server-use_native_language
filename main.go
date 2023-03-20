package main

import (
	"afterclass/server"
	"afterclass/server/router"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	defer func() {
		if data := recover(); data != nil {
			log.Fatal(data)
		}
		fmt.Println("Program is ending......")
	}()
	shutDown := server.NewGracefulShutDown()

	servers := server.NewServer("Web",
		server.ComputeTimeBuilder,
		server.SayByeBuilder,
		shutDown.ShutDownFilterBuilder)

	//注册路由
	servers.Router("GET", "/home", router.Home)
	servers.Router("GET", "/sign", router.Sign)

	//启动服务
	go func() {
		err := servers.Start(":8080")
		if err != nil {
			panic(err)
		}
	}()

	server.WaitShutDone(
		func(ctx context.Context) error {
			fmt.Println("mock notify gateway")
			time.Sleep(time.Second * 2)
			return nil
		},
		server.BuildCloseServerHook(servers),
	)
}

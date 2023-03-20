package server

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

/*
os.Exit()强制退出
*/
//等待用户的信号输入来停止服务
func WaitShutDone() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, ShutdownSignals...)
	select {
	//当用户信号传来时
	case sig := <-signals:
		fmt.Printf("recv signal from user :%v,and shuting down the server.....\n", sig)
		//若10分钟后还没退出则强制退出，因为中间要执行狗子函数
		time.AfterFunc(time.Minute*10, func() {
			fmt.Println("time out!!10 minute!! application shut down!!")
			os.Exit(1)
		})
		//TODO hook函数
		os.Exit(1)
	}
}

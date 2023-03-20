package server

import (
	"afterclass/server/router"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

var ErrorHookTimeout = errors.New("hook time out")

//实现请求拦截
type GracefulShutdown struct {
	//还在处理中的请求数量
	reqCnt int64
	//大于0 表示要关闭了
	closing int32

	//传递已经处理完所有请求的信号
	zeroReqCnt chan struct{}
}

func NewGracefulShutDown() *GracefulShutdown {
	return &GracefulShutdown{
		zeroReqCnt: make(chan struct{}),
	}
}

func (g *GracefulShutdown) ShutDownFilterBuilder(next Filter) Filter {
	return func(ctx *router.MyContext) {
		//开始拒绝所有请求,即closing > 1
		cl := atomic.LoadInt32(&g.closing)
		if cl > 0 {
			ctx.W.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		//将该请求入列
		atomic.AddInt64(&g.reqCnt, 1)
		next(ctx)
		//请求处理完就出列
		n := atomic.AddInt64(&g.reqCnt, -1)
		if cl > 0 && n == 0 {
			g.zeroReqCnt <- struct{}{}
		}
	}
}

//这是一个hook
func (g *GracefulShutdown) RejectNewReqAndWait(ctx context.Context) error {
	//由该函数来发起拒绝请求的信号
	atomic.AddInt32(&g.closing, 1)

	if atomic.LoadInt64(&g.reqCnt) == 0 {
		return nil
	}
	done := ctx.Done()
	select {
	case <-done:
		fmt.Printf("time out \n")
		return ErrorHookTimeout
	case <-g.zeroReqCnt:
		fmt.Println("请求全部处理完了")
	}
	return nil
}

/*
os.Exit()强制退出
*/
//等待用户的信号输入来停止服务
func WaitShutDone(hook ...Hook) {
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
		for _, h := range hook {
			//执行hook函数，每个hook设置30s的执行时间
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			err := h(ctx)
			if err != nil {
				fmt.Printf("failed to run hook :%v\n", err)
			}
			cancel()
		}
		os.Exit(0)
	}
}

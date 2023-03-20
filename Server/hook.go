package server

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Hook func(ctx context.Context) error

func BuildCloseServerHook(servers ...Server) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		doneCh := make(chan struct{})
		wg.Add(len(servers))

		//单独关闭每一个server
		for _, s := range servers {
			go func(svr Server) {
				err := svr.ShutDown(ctx)
				if err != nil {
					fmt.Printf(" server shut down failed:%v\n", err)
				}
				time.Sleep(time.Second)
				wg.Done()
			}(s)
		}

		go func() {
			wg.Wait()
			//当所有服务都完成关闭时传入信号
			doneCh <- struct{}{}
		}()

		//当server关不掉时，利用ctx超时关闭
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return errors.New("server close time out")
		case <-doneCh:
			fmt.Println("all server have closed")
			return nil
		}
	}
}

package server

import (
	"afterclass/server/router"
	"fmt"
	"time"
)

/*
实现aop 来封装我们的请求
filte 过滤
*/

type FilterBuilder func(next Filter) Filter

type Filter func(ctx *router.MyContext)

//计算时间戳
func ComputeTimeBuilder(next Filter) Filter {
	return func(ctx *router.MyContext) {
		start := time.Now().Nanosecond()
		next(ctx)
		fmt.Printf("take time :%d\n", time.Now().Nanosecond()-start)
	}
}

// 打招呼过滤器
func SayByeBuilder(next Filter) Filter {
	return func(ctx *router.MyContext) {
		next(ctx)
		fmt.Println("bye~~~")
	}
}

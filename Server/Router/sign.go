package router

import (
	"fmt"
	"net/http"
)

type UserData struct {
	Name string `json:"name"`
	Age  int    `json:"ag"`
}

type RespData struct {
	Data string `json:"data"`
}

func Sign(ctx *MyContext) {
	//用户数据会以json的形式塞在body中传送过来
	u_data := &UserData{}
	err := ctx.ParseJson(u_data)
	if err != nil {
		fmt.Fprintf(ctx.W, "parse json failed")
	}
	//装数据并返回
	respdata := &RespData{Data: "123"}
	err = ctx.RespJson(http.StatusOK, respdata)
	if err != nil {
		fmt.Printf("send json failed:%v\n", err)
		return
	}
}

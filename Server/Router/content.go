package router

import (
	"encoding/json"
	"io"
	"net/http"
)

type MyContext struct {
	W http.ResponseWriter
	R *http.Request
}

func CreateCtx(w http.ResponseWriter, r *http.Request) *MyContext {
	return &MyContext{W: w, R: r}
}

func (ctx *MyContext) ParseJson(obj interface{}) error {
	body, err := io.ReadAll(ctx.R.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, obj)
	//parse json
	if err != nil {
		return err
	}
	return nil
}

func (ctx *MyContext) RespJson(code int, obj interface{}) error {
	resp, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	ctx.W.WriteHeader(code)
	_, err = ctx.W.Write(resp)
	return err
}

func (ctx *MyContext) OkJson(obj interface{}) error {
	return ctx.RespJson(http.StatusOK, obj)
}

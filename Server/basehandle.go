package server

import (
	"afterclass/server/router"
	"net/http"
)

type BaseHandle struct {
	M map[string]func(ctx *router.MyContext)
}

func (h *BaseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//读取请求路径和方法，并进行拼接
	key := h.Key(r.Method, r.URL.Path)
	if handler, ok := h.M[key]; ok {
		//存在就执行这个方法
		handler(router.CreateCtx(w, r))
	} else {
		//不存在返回404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("page not found"))
	}
}

func (h *BaseHandle) Key(method string, url string) string {
	return method + "#" + url
}

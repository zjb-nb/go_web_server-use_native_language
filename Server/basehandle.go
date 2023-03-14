package server

import (
	"afterclass/server/router"
	"net/http"
)

type RouteBle interface {
	Router(method string, url string, handleFunc MyHandleFunc)
}

type BaseHandle interface {
	http.Handler
	RouteBle
}

// 不想将m暴露给外部，那我就自己控制m的调用
type BaseHandleOnMap struct {
	m map[string]func(ctx *router.MyContext)
}

func (h *BaseHandleOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//读取请求路径和方法，并进行拼接
	key := h.Key(r.Method, r.URL.Path)
	if handler, ok := h.m[key]; ok {
		//存在就执行这个方法
		handler(router.CreateCtx(w, r))
	} else {
		//不存在返回404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("page not found"))
	}
}

func (h *BaseHandleOnMap) Router(method string, url string, handleFunc MyHandleFunc) {
	key := h.Key(method, url)
	h.m[key] = handleFunc
}

func (h *BaseHandleOnMap) Key(method string, url string) string {
	return method + "#" + url
}

func NewBaseHandleOnMap() BaseHandle {
	return &BaseHandleOnMap{
		m: make(map[string]func(ctx *router.MyContext), 0),
	}
}

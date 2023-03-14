package server

import (
	"afterclass/server/router"
	"net/http"
)

type MyHandleFunc func(ctx *router.MyContext)

type Server interface {
	RouteBle
	Start(port string) error
}

type SDKServer struct {
	Name    string
	handler BaseHandle
	//假设filter全局挂载
	root Filter
}

func (s *SDKServer) Router(method string, url string, handleFunc MyHandleFunc) {
	s.handler.Router(method, url, handleFunc)
}

func (s *SDKServer) Start(port string) error {
	//挂载，我们需要让每个请求都执行过滤器 root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := router.CreateCtx(w, r)
		s.root(ctx)
	})

	return http.ListenAndServe(port, nil)
}

func NewServer(str string, builders ...FilterBuilder) Server {
	handler := NewBaseHandleOnMap()
	//执行请求的filter为最后一个
	var root Filter = func(ctx *router.MyContext) {
		handler.ServeHTTP(ctx.W, ctx.R)
	}

	for i := len(builders) - 1; i >= 0; i-- {
		buildFunc := builders[i]
		root = buildFunc(root)
	}

	return &SDKServer{
		Name:    str,
		handler: handler,
		root:    root,
	}
}

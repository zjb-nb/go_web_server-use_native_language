package server

import (
	"afterclass/server/router"
	"net/http"
)

type Server interface {
	Router(url string, handler func(ctx *router.MyContext))
	Start(port string) error
}

type SDKServer struct {
	Name string
}

func (s *SDKServer) Router(url string, handler func(ctx *router.MyContext)) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		//在这里我们生成ctx
		ctx := router.CreateCtx(w, r)
		handler(ctx)
	})
}

func (s *SDKServer) Start(port string) error {
	return http.ListenAndServe(port, nil)
}

func NewServer(str string) Server {
	return &SDKServer{Name: str}
}

package server

import (
	"afterclass/server/router"
	"net/http"
)

type Server interface {
	Router(method string, url string, handleFunc func(ctx *router.MyContext))
	Start(port string) error
}

type SDKServer struct {
	Name    string
	handler *BaseHandle
}

func (s *SDKServer) Router(method string, url string, handleFunc func(ctx *router.MyContext)) {
	key := s.handler.Key(method, url)
	s.handler.M[key] = handleFunc
}

func (s *SDKServer) Start(port string) error {
	//挂载
	http.Handle("/", s.handler)
	return http.ListenAndServe(port, nil)
}

func NewServer(str string) Server {
	return &SDKServer{Name: str, handler: &BaseHandle{M: make(map[string]func(ctx *router.MyContext))}}
}

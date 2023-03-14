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
}

func (s *SDKServer) Router(method string, url string, handleFunc MyHandleFunc ) {
	s.handler.Router(method, url, handleFunc)
}

func (s *SDKServer) Start(port string) error {
	//挂载
	http.Handle("/", s.handler)
	return http.ListenAndServe(port, nil)
}

func NewServer(str string) Server {
	return &SDKServer{
		Name:    str,
		handler: NewBaseHandleOnMap(),
	}
}

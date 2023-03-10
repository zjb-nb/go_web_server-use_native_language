package server

import "net/http"

type Server interface {
	Router(url string, handler http.HandlerFunc)
	Start(port string) error
}

type SDKServer struct {
	Name string
}

func (s *SDKServer) Router(url string, handler http.HandlerFunc) {
	http.HandleFunc(url, handler)
}

func (s *SDKServer) Start(port string) error {
	return http.ListenAndServe(port, nil)
}

func NewServer(str string) Server {
	return &SDKServer{Name: str}
}

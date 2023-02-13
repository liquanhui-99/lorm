package lorm

import (
	"net"
	"net/http"
)

type HandleFunc func()

type Server interface {
	http.Handler
	Start(addr string) error
	AddRouter(method string, path string, handle HandleFunc)
}

type HTTPServer struct {
	*Router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Router: newRouter(),
	}
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("HTTPServer")
}

func (h *HTTPServer) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return http.Serve(listener, h)
}

func (h *HTTPServer) AddRouter(method string, path string, handle HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPServer) GET(path string, handle HandleFunc) {
	h.AddRouter(http.MethodGet, path, handle)
}

func (h *HTTPServer) POST(path string, handle HandleFunc) {
	h.AddRouter(http.MethodPost, path, handle)
}

func (h *HTTPServer) PUT(path string, handle HandleFunc) {
	h.AddRouter(http.MethodPut, path, handle)
}

func (h *HTTPServer) PATCH(path string, handle HandleFunc) {
	h.AddRouter(http.MethodPatch, path, handle)
}

func (h *HTTPServer) DELETE(path string, handle HandleFunc) {
	h.AddRouter(http.MethodDelete, path, handle)
}

package server

import (
	"HttpServer/src/packages/server/handlers"
	"net/http"
)

type MyServer struct {
	*http.Server
}

func (s *MyServer) StartServer() error {
	return http.ListenAndServe(":5252", GetRouter(&handlers.MyHandler{}))
}

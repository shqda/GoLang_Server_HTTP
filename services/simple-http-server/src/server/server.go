package server

import (
	"HttpServer/config"
	"HttpServer/server/handlers"
	"net/http"
	"strconv"
)

type MyServer struct {
	*http.Server
}

func (s *MyServer) StartServer(cfg *config.ServerConfig) error {
	return http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), GetRouter(&handlers.MyHandler{}))
}

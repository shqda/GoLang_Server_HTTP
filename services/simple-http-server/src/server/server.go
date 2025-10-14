package server

import (
	"HttpServer/config"
	"HttpServer/server/handlers"
	"net/http"
)

type MyServer struct {
	*http.Server
}

func (s *MyServer) StartServer(cfg *config.ServerConfig) error {
	return http.ListenAndServe(cfg.Server.Host+":"+cfg.GetServerPortAsString(), GetRouter(&handlers.MyHandler{}))
}

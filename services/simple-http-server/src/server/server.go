package server

import (
	"HttpServer/config"
	"HttpServer/server/handlers"
	"fmt"
	"net/http"
	"strconv"
)

type MyServer struct {
	*http.Server
}

func (s *MyServer) StartServer(config *config.ServerConfig) error {
	fmt.Printf("Starting server on port %v", config.Server.Port)
	return http.ListenAndServe(config.Server.Host+":"+strconv.Itoa(config.Server.Port), GetRouter(&handlers.MyHandler{}))
}

package server

import (
	"HttpServer/configs"
	"HttpServer/server/handlers"
	"fmt"
	"net/http"
	"strconv"
)

type MyServer struct {
	*http.Server
}

func (s *MyServer) StartServer(config *configs.ServerConfig) error {
	fmt.Printf("Starting server on port %v", config.Server.Port)
	return http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), GetRouter(&handlers.MyHandler{}))
}
